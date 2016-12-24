package main

import (
	"fmt"

	"github.com/xrash/smetrics"
)

type distanceMatrix struct {
	links  []string
	matrix [][]float64
	size   int
}

func newMatrix(links []string) *distanceMatrix {
	linkcount := len(links)

	// 빈 매트릭스 일단 만들기
	// 전부 0으로 채워둔다
	m := make([][]float64, linkcount)
	for i := 0; i < linkcount; i++ {
		m[i] = make([]float64, linkcount)
	}

	// 대각선은 1로. 자기자신이 들어가니까
	for i := 0; i < linkcount; i++ {
		m[i][i] = 1
	}

	for i, link := range links {
		for j := i + 1; j < linkcount; j++ {
			otherlink := links[j]
			dist := smetrics.Jaro(link, otherlink)
			m[i][j] = dist
			m[j][i] = dist
		}
	}

	return &distanceMatrix{
		links:  links,
		matrix: m,
		size:   linkcount,
	}
}

type distanceCluster struct {
	links       []string
	clusterlist []int
	allowed     float64
}

// 유사도가 비슷한 문자열끼리 같은 클러스터 ID를 부여한다
func newCluster(links []string, allowed float64) *distanceCluster {
	m := newMatrix(links)

	clusterid := 0
	clusterlist := make([]int, m.size)
	for i := 0; i < m.size; i++ {
		clusterlist[i] = -1
	}

	for i := 0; i < m.size; i++ {
		if clusterlist[i] >= 0 {
			continue
		}
		currcluster := clusterid
		clusterid++

		clusterlist[i] = currcluster
		for j := i + 1; j < m.size; j++ {
			if m.matrix[i][j] >= allowed {
				clusterlist[j] = currcluster
			} else {
				// 불연속구간은 같은 클러스터로 취급하지 않는다
				break
			}
		}
	}

	return &distanceCluster{
		links:       links,
		clusterlist: clusterlist,
		allowed:     allowed,
	}
}

func (c *distanceCluster) MaxClusterID() int {
	last := len(c.links) - 1
	return c.clusterlist[last]
}
func (c *distanceCluster) Show() {
	for i, link := range c.links {
		cid := c.clusterlist[i]
		fmt.Printf("cid=%d : %s\n", cid, link)
	}
}

func (c *distanceCluster) GetCluster(id int) []string {
	links := []string{}
	for i, cid := range c.clusterlist {
		if cid == id {
			links = append(links, c.links[i])
		}
	}
	return links
}
