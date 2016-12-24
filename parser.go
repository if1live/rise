package main

/*
https://schier.co/blog/2015/04/26/a-simple-web-scraper-in-go.html
*/

import (
	"io/ioutil"
	"net/http"
	"strings"

	"net/url"

	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
	"golang.org/x/net/html"
)

func makeSafeURI(uri string) string {
	// http:// 없을때 기본값 넣어주기
	parsed, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	if parsed.Scheme == "" {
		return "http://" + uri
	}
	return uri
}

func getHTMLText(uri string) string {
	cachedir := "./_cache"
	cache := diskcache.New(cachedir)
	tp := httpcache.NewTransport(cache)
	client := &http.Client{Transport: tp}

	uri = makeSafeURI(uri)
	resp, err := client.Get(uri)
	if err != nil {
		panic(err)
	}

	bytes, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return string(bytes)
}

func getLink(t html.Token) string {
	for _, a := range t.Attr {
		if a.Key == "src" {
			return a.Val
		}
	}
	return ""
}

func parseHTMLText(htmltext string) []string {
	r := strings.NewReader(htmltext)
	z := html.NewTokenizer(r)

	links := []string{}

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return links
		case tt == html.SelfClosingTagToken:
			t := z.Token()
			isImg := (t.Data == "img")
			if isImg {
				link := getLink(t)
				if link != "" {
					links = append(links, link)
				}
			}
		case tt == html.StartTagToken:
			t := z.Token()
			isImg := (t.Data == "img")
			if isImg {
				link := getLink(t)
				if link != "" {
					links = append(links, link)
				}
			}
		}
	}
}
