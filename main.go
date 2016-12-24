package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"flag"

	"log"

	"path/filepath"

	"github.com/gregjones/httpcache"
	"github.com/if1live/staticfilecache"
	"github.com/jhoonb/archivex"
)

var cmd string
var uri string
var cid int
var similarity float64
var workerCount int
var filter int
var output string

type fetchConfig struct {
	URL        string  `json:"url"`
	ClusterID  int     `json:"cluster_id"`
	Filter     int     `json:"filter"`
	Similarity float64 `json:"similarity"`
}

func (m *fetchConfig) Marshal() []byte {
	data, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	var out bytes.Buffer
	json.Indent(&out, data, "", "  ")
	return out.Bytes()
}

func makeClusterFromFlag() *distanceCluster {
	htmltext := getHTMLText(uri)
	links := parseHTMLText(htmltext)
	for i, link := range links {
		links[i] = ApplyFilter(link, filter)
	}
	return newCluster(links, similarity)
}

func init() {
	flag.StringVar(&cmd, "cmd", "help", "command")
	flag.StringVar(&uri, "uri", "", "target uri")
	flag.IntVar(&cid, "cid", -1, "cluster id")
	flag.Float64Var(&similarity, "similarity", 0.9, "similarity")
	flag.IntVar(&workerCount, "worker", 8, "worker count")
	flag.IntVar(&filter, "filter", 0, "link filter")
	flag.StringVar(&output, "output", "download.zip", "output zip filename")
}

func main() {
	flag.Parse()

	switch cmd {
	case "show":
		mainForShow()
	case "fetch":
		mainForFetch()
	default:
		mainForHelp()
	}
}

func mainForShow() {
	cluster := makeClusterFromFlag()
	if cid < 0 || cid > cluster.MaxClusterID() {
		cluster.Show()
	} else {
		links := cluster.GetCluster(cid)
		for i, link := range links {
			fmt.Printf("(%d/%d) %s\n", i+1, len(links), link)
		}
	}
}

func mainForHelp() {
	fmt.Println("help")
}

type fetchCommand struct {
	uri string
	idx int
}
type fetchResult struct {
	resp *http.Response
	uri  string
	idx  int
}

func workerFetch(id int, jobs <-chan *fetchCommand, results chan<- *fetchResult) {
	for req := range jobs {
		if req.uri == "" {
			return
		}
		uri := req.uri

		log.Printf("[worker %d] download %s\n", id, uri)
		cachedir := "./_cache_static"
		cache := staticfilecache.New(cachedir)
		tp := httpcache.NewTransport(cache)
		client := &http.Client{Transport: tp}
		resp, _ := client.Get(uri)

		results <- &fetchResult{
			resp: resp,
			uri:  uri,
			idx:  req.idx,
		}
	}
}

func mainForFetch() {
	cluster := makeClusterFromFlag()
	links := cluster.GetCluster(cid)
	linkcount := len(links)
	if linkcount == 0 {
		log.Printf("No link found, cid=%d\n", cid)
		return
	}

	jobs := make(chan *fetchCommand, linkcount)
	fetchResults := make(chan *fetchResult, linkcount)
	zipFinishCh := make(chan bool, linkcount)

	zip := new(archivex.ZipFile)
	zip.Create(output)

	// 다운로드에 사용한 정보 저장하기
	config := &fetchConfig{
		URL:        uri,
		ClusterID:  cid,
		Filter:     filter,
		Similarity: similarity,
	}
	zip.Add("metadata.json", config.Marshal())

	go func(z *archivex.ZipFile, count int, results chan *fetchResult, finishCh chan bool) {
		for i := 1; i <= count; i++ {
			result := <-results
			log.Printf("complete (%d/%d)\n", i, linkcount)
			bytes, _ := ioutil.ReadAll(result.resp.Body)
			result.resp.Body.Close()

			basename := filepath.Base(result.uri)
			filename := MakeFilename(basename, result.idx)

			z.Add(filename, bytes)
		}

		finishCh <- true
	}(zip, linkcount, fetchResults, zipFinishCh)

	for w := 1; w <= workerCount; w++ {
		go workerFetch(w, jobs, fetchResults)
	}
	for i, link := range links {
		jobs <- &fetchCommand{
			uri: link,
			idx: i,
		}
	}
	// 종료 지점 넣어주기. 요청을 받은 fetch worker는 알아서 종료될것이다
	for i := 0; i < linkcount; i++ {
		jobs <- &fetchCommand{
			uri: "",
			idx: -1,
		}
	}
	// zip 파일쓰기 끝난거 확인하면 종료
	<-zipFinishCh
	zip.Close()
}
