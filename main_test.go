package main

import (
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func ReadLinksFromFile(filepath string) []string {
	bytes, _ := ioutil.ReadFile(filepath)
	text := string(bytes)
	links := strings.Split(text, "\n")
	for i, link := range links {
		links[i] = strings.Replace(link, "\r", "", -1)
	}
	return links
}

func TestMain(t *testing.T) {
	cases := []struct {
		htmlfile     string
		clusterids   []int
		similarity   float64
		filter       int
		expectedfile string
	}{
		{"testdata/001.html", []int{4}, 0.85, 1, "testdata/001.txt"},
	}

	for _, c := range cases {
		htmlbytes, _ := ioutil.ReadFile(c.htmlfile)
		htmltext := string(htmlbytes)
		links := parseHTMLText(htmltext)
		for i, link := range links {
			links[i] = ApplyFilter(link, c.filter)
		}

		cluster := newClusterList(links, c.similarity)
		actuallinks := cluster.GetClusters(c.clusterids)

		expectedlinks := ReadLinksFromFile(c.expectedfile)

		if !reflect.DeepEqual(actuallinks, expectedlinks) {
			t.Error(
				"For Main",
				"expected", expectedlinks,
				"got", actuallinks,
			)
		}
	}
}
