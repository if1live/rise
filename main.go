package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"flag"

	"log"

	"path/filepath"

	"strconv"

	"github.com/gregjones/httpcache"
	"github.com/if1live/staticfilecache"
	"github.com/jhoonb/archivex"
)

var cmd string
var uri string
var cid string
var similarity float64
var workerCount int
var filter int
var output string

func clusterIdList() []int {
	cids := []int{}
	for _, c := range strings.Split(cid, ",") {
		clusterid, err := strconv.Atoi(c)
		if err != nil {
			continue
		}
		cids = append(cids, clusterid)
	}
	return cids
}

type fetchConfig struct {
	URL           string  `json:"url"`
	ClusterIDList []int   `json:"cluster_id_list"`
	Filter        int     `json:"filter"`
	Similarity    float64 `json:"similarity"`
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

func makeClusterFromFlag() *clusterList {
	htmltext := getHTMLText(uri)
	links := parseHTMLText(htmltext)
	for i, link := range links {
		links[i] = ApplyFilter(link, filter)
	}
	return newClusterList(links, similarity)
}

func init() {
	flag.StringVar(&cmd, "cmd", "help", "command")
	flag.StringVar(&uri, "uri", "", "target uri")
	flag.StringVar(&cid, "cid", "", "cluster id")
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
	if cid == "" {
		cluster.Show()
	} else {
		links := cluster.GetClusters(clusterIdList())
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

		imguri := req.uri
		imguri = makeSafeImageURI(imguri, uri)

		log.Printf("[worker %d] download %s\n", id, imguri)
		cachedir := "./_cache_static"
		cache := staticfilecache.New(cachedir)
		tp := httpcache.NewTransport(cache)
		client := &http.Client{Transport: tp}
		resp, _ := client.Get(imguri)

		results <- &fetchResult{
			resp: resp,
			uri:  imguri,
			idx:  req.idx,
		}
	}
}

func mainForFetch() {
	cluster := makeClusterFromFlag()
	cids := clusterIdList()
	links := cluster.GetClusters(cids)

	linkcount := len(links)
	if linkcount == 0 {
		log.Printf("No link found, cid=%s\n", cid)
		return
	}

	jobs := make(chan *fetchCommand, linkcount)
	fetchResults := make(chan *fetchResult, linkcount)
	zipFinishCh := make(chan bool, linkcount)

	zip := new(archivex.ZipFile)
	zip.Create(output)

	// 다운로드에 사용한 정보 저장하기
	config := &fetchConfig{
		URL:           uri,
		ClusterIDList: cids,
		Filter:        filter,
		Similarity:    similarity,
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
