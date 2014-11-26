package graphs

import (
	"algs/symboltables"
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
	epq    *symboltables.RBT
	Source string
}

func (spf *SPF) Init(g *Graph) {
	spf.graph = g
	spf.spt = make(map[string]string)
	spf.distTo = make(map[string]int32)
	spf.epq = new(symboltables.RBT)
	for _, v := range (spf.graph).Vertices() {
		spf.distTo[v] = SPF_INFINITY
	}
}

func (spf *SPF) SP(s string) {
	spf.Source = s
	spf.distTo[s] = 0
	spf.epq.Put(Edge{to: s, dist: 0})
	for spf.epq.Len() > 0 {
		spf.relax()
	}
}

func (spf *SPF) relax() {
	Min := spf.epq.FindMin()
	min := Min.(Edge).to
	spf.epq.DeleteMin()
	for _, adj := range spf.graph.Adjacency(min) {
		weight := spf.graph.adjacency[min][adj]
		if spf.distTo[adj] > spf.distTo[min]+weight {
			spf.distTo[adj] = spf.distTo[min] + weight
			spf.spt[adj] = min
			spf.epq.Put(Edge{to: adj, dist: spf.distTo[adj]})
		}
	}
}

func (spf *SPF) SPFDist() map[string]int32 {
	return spf.distTo
}

func (spf *SPF) SPFPath() map[string]string {
	return spf.spt
}
