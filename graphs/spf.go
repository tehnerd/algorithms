package graphs

const SPF_INFINITY = 16000000

type SPF struct {
	graph  *Graph
	spt    map[string]string
	distTo map[string]int32
	pq     map[string]int32 //TODO: add RBT instead of map
	Source string
}

func (spf *SPF) Init(g *Graph) {
	spf.graph = g
	spf.spt = make(map[string]string)
	spf.pq = make(map[string]int32)
	spf.distTo = make(map[string]int32)
	for _, v := range (spf.graph).Vertices() {
		spf.distTo[v] = SPF_INFINITY
	}
}

func (spf *SPF) SP(s string) {
	spf.Source = s
	spf.distTo[s] = 0
	spf.pq[s] = 0
	for len(spf.pq) > 0 {
		spf.relax()
	}
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
