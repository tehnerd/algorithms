package graphs

import (
	"algs/heap"
	"fmt"
	"time"
)

const SPF_INFINITY = 16000000

type Edge struct {
	to   string
	dist int32
}

type EdgesPQ struct {
	edges []Edge
}

func (epq *EdgesPQ) Len() int {
	return len(epq.edges)
}

func (epq *EdgesPQ) Compare(i, j int) int {
	if epq.edges[i].dist < epq.edges[j].dist {
		return -1
	} else if epq.edges[i].dist > epq.edges[j].dist {
		return 1
	} else {
		return 0
	}
}

func (epq *EdgesPQ) Swap(i, j int) {
	tmp := epq.edges[i]
	epq.edges[i] = epq.edges[j]
	epq.edges[j] = tmp
}

func (epq *EdgesPQ) Insert(e Edge) {
	epq.edges = append(epq.edges, e)
	heap.MinHeapCheck(epq, len(epq.edges)-1)
}

func (epq *EdgesPQ) Remove() Edge {
	edge := epq.edges[0]
	epq.edges[0] = epq.edges[len(epq.edges)-1]
	epq.edges = epq.edges[:len(epq.edges)-1]
	heap.MinHeapCheck(epq, 0)
	return edge
}

type SPF struct {
	graph  *Graph
	spt    map[string]string
	distTo map[string]int32
	pq     map[string]int32 //TODO: add RBT instead of map
	epq    *EdgesPQ
	Source string
}

func (spf *SPF) Init(g *Graph) {
	spf.graph = g
	spf.spt = make(map[string]string)
	spf.pq = make(map[string]int32)
	spf.epq = new(EdgesPQ)
	spf.epq.edges = make([]Edge, 0)
	spf.distTo = make(map[string]int32)
	for _, v := range (spf.graph).Vertices() {
		spf.distTo[v] = SPF_INFINITY
	}
}

func (spf *SPF) SP(s string) {
	spf.Source = s
	spf.distTo[s] = 0
	spf.pq[s] = 0
	t1 := time.Now()
	for len(spf.pq) > 0 {
		spf.relax()
	}
	t2 := time.Now()
	fmt.Println(t2.UnixNano() - t1.UnixNano())
}

func (spf *SPF) relax() {
	min := "none"
	minVal := int32(-1)
	for k, v := range spf.pq {
		if (minVal == -1) || (v <= minVal) {
			minVal = v
			min = k
		}
	}
	delete(spf.pq, min)
	for _, adj := range spf.graph.Adjacency(min) {
		weight := spf.graph.adjacency[min][adj]
		if spf.distTo[adj] > spf.distTo[min]+weight {
			spf.distTo[adj] = spf.distTo[min] + weight
			spf.spt[adj] = min
			if _, exist := spf.pq[adj]; exist {
				spf.pq[adj] = spf.distTo[adj]
			} else {
				spf.pq[adj] = spf.distTo[adj]
			}
		}
	}
}

func (spf *SPF) SPFDist() map[string]int32 {
	return spf.distTo
}

func (spf *SPF) SPFPath() map[string]string {
	return spf.spt
}
