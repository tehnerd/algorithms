package graphs

import (
	"algs/symboltables"
	"fmt"
	"reflect"
	"time"
)

const SPF_INFINITY = 16000000

type Edge struct {
	to   string
	dist int32
}

func (e1 Edge) Compare(e2ckv symboltables.ComparableKV) int {
	switch e2 := e2ckv.(type) {
	case Edge:
		if e1.dist < e2.dist {
			return -1
		} else if e1.dist > e2.dist {
			return 1
		} else if e1.to < e2.to {
			return -1
		} else if e1.to < e2.to {
			return 1
		} else {
			return 0
		}
	default:
		return -100
	}
}

type SPF struct {
	graph  *Graph
	spt    map[string]string
	distTo map[string]int32
	pq     map[string]int32 //TODO: add RBT instead of map
	epq    *symboltables.RBT
	Source string
}

func (spf *SPF) Init(g *Graph) {
	spf.graph = g
	spf.spt = make(map[string]string)
	spf.pq = make(map[string]int32)
	spf.distTo = make(map[string]int32)
	spf.epq = new(symboltables.RBT)
	for _, v := range (spf.graph).Vertices() {
		spf.distTo[v] = SPF_INFINITY
	}
}

func (spf *SPF) SP(s string) {
	spf.Source = s
	spf.distTo[s] = 0
	spf.pq[s] = 0
	spf.epq.Put(Edge{to: s, dist: 0})
	/*
		t1 := time.Now()
		for len(spf.pq) > 0 {
			spf.relax()
		}
	*/
	t1 := time.Now()
	for spf.epq.Len() > 0 {
		spf.relax()
	}
	t2 := time.Now()
	fmt.Println(t2.UnixNano() - t1.UnixNano())
}

func (spf *SPF) relax() {
	/*
		min := "none"
		minVal := int32(-1)
		for k, v := range spf.pq {
			if (minVal == -1) || (v <= minVal) {
				minVal = v
				min = k
			}
		}
	*/
	Min := spf.epq.FindMin()
	min := reflect.ValueOf(Min).Interface().(Edge).to
	spf.epq.DeleteMin()
	//delete(spf.pq, min)
	for _, adj := range spf.graph.Adjacency(min) {
		weight := spf.graph.adjacency[min][adj]
		if spf.distTo[adj] > spf.distTo[min]+weight {
			spf.distTo[adj] = spf.distTo[min] + weight
			spf.spt[adj] = min
			spf.epq.Put(Edge{to: adj, dist: spf.distTo[adj]})
			/*
				if _, exist := spf.pq[adj]; exist {
					spf.pq[adj] = spf.distTo[adj]
				} else {
					spf.pq[adj] = spf.distTo[adj]
				}
			*/
		}
	}
}

func (spf *SPF) SPFDist() map[string]int32 {
	return spf.distTo
}

func (spf *SPF) SPFPath() map[string]string {
	return spf.spt
}
