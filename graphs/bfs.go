package graphs

type BreadthFirstPath struct {
	marked map[string]bool
	edgeTo map[string]string
	source string
}

func (bfp *BreadthFirstPath) Init(g *Graph, s string) {
	if bfp.marked == nil {
		bfp.marked = make(map[string]bool)
		bfp.edgeTo = make(map[string]string)
		bfp.source = s
	}
	for _, v := range g.Vertices() {
		bfp.marked[v] = false
		bfp.edgeTo[v] = ""
	}
}

type stringQueue struct {
	val  string
	next *stringQueue
}

func (sq *stringQueue) enqueue(s string) {
	for sq.next != nil {
		sq = sq.next
	}
	if sq.val != "" {
		sq.next = &stringQueue{val: s}
	} else {
		sq.val = s
	}
}

func (sq *stringQueue) dequeue() string {
	sqPrev := new(stringQueue)
	for sq.next != nil {
		sqPrev = sq
		sq = sq.next
	}
	sqPrev.next = nil
	if sqPrev.val == "" {
		firstVal := sq.val
		sq.val = ""
		return firstVal
	}
	return sq.val
}

func (sq *stringQueue) isEmpty() bool {
	return sq.val == ""
}

func (bfp *BreadthFirstPath) BFP(g *Graph, s string) {
	sq := new(stringQueue)
	bfp.Init(g, s)
	bfp.marked[s] = true
	sq.enqueue(s)
	for !sq.isEmpty() {
		v := sq.dequeue()
		for _, av := range g.Adjacency(v) {
			if !bfp.marked[av] {
				bfp.edgeTo[av] = v
				bfp.marked[av] = true
				sq.enqueue(av)
			}
		}
	}
}

func (bfp *BreadthFirstPath) BFPHasPathTo(g *Graph, s, d string) bool {
	bfp.BFP(g, s)
	return bfp.marked[d]
}

func (bfp *BreadthFirstPath) BFPPathTo(g *Graph, s, d string) []string {
	path := make([]string, 0)
	bfp.BFP(g, s)
	if !bfp.marked[d] {
		return path
	}
	for v := d; v != s; v = bfp.edgeTo[v] {
		path = append(path, v)
	}
	path = append(path, s)
	return path

}
