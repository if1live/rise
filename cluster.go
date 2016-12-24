package main

import (
	"fmt"

	"github.com/xrash/smetrics"
)

type clusterList struct {
	words       []string
	clusterlist []int
	similarity  float64
}

// 연속된 단어-단어의 유사도만으로 클러스터를 구성해도 문제없다
func newClusterList(words []string, similarity float64) *clusterList {
	if len(words) <= 1 {
		return &clusterList{}
	}

	wordcount := len(words)
	clusterlist := make([]int, wordcount)
	for i := 0; i < wordcount; i++ {
		clusterlist[i] = -1
	}

	clusterid := 0
	clusterlist[0] = clusterid
	for i := 0; i < wordcount-1; i++ {
		w1 := words[i]
		w2 := words[i+1]

		dist := smetrics.Jaro(w1, w2)
		if dist >= similarity {
			clusterlist[i+1] = clusterid
		} else {
			clusterid++
			clusterlist[i+1] = clusterid
		}
	}
	return &clusterList{
		words:       words,
		clusterlist: clusterlist,
		similarity:  similarity,
	}
}

func (c *clusterList) MaxClusterID() int {
	last := len(c.words) - 1
	return c.clusterlist[last]
}

func (c *clusterList) Show() {
	for i, word := range c.words {
		cid := c.clusterlist[i]
		fmt.Printf("cid=%d : %s\n", cid, word)
	}
}

func (c *clusterList) GetCluster(id int) []string {
	words := []string{}
	for i, cid := range c.clusterlist {
		if cid == id {
			words = append(words, c.words[i])
		}
	}
	return words
}

func (c *clusterList) GetClusters(ids []int) []string {
	words := []string{}
	for _, cid := range ids {
		words = append(words, c.GetCluster(cid)...)
	}
	return words
}
