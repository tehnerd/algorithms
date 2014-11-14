package graphs

type DepthFirstPath struct {
	marked map[string]bool
	edgeTo map[string]string
	source string
}

func (dfp *DepthFirstPath) Init(g *Graph, s string) {
	if dfp.marked == nil {
		dfp.marked = make(map[string]bool)
		dfp.edgeTo = make(map[string]string)
		dfp.source = s
	}
	for _, v := range g.Vertices() {
		dfp.marked[v] = false
		dfp.edgeTo[v] = ""
	}

}

func (dfp *DepthFirstPath) DFP(g *Graph, s string) {
	dfp.marked[s] = true
	for _, av := range g.Adjacency(s) {
		if !dfp.marked[av] {
			dfp.edgeTo[av] = s
			dfp.DFP(g, av)
		}
	}
}

func (dfp *DepthFirstPath) hasPathTo(d string) bool {
	return dfp.marked[d]
}

func createDFPContext(g *Graph, s string) *DepthFirstPath {
	dfpContext := new(DepthFirstPath)
	dfpContext.Init(g, s)
	dfpContext.DFP(g, s)
	return dfpContext
}

func HastPathTo(g *Graph, s, d string) bool {
	dfpContext := createDFPContext(g, s)
	return dfpContext.hasPathTo(d)
}

func PathTo(g *Graph, s, d string) []string {
	path := make([]string, 0)
	dfpContext := createDFPContext(g, s)
	if !dfpContext.hasPathTo(d) {
		return path
	}
	for v := d; v != s; v = dfpContext.edgeTo[v] {
		path = append(path, v)
	}
	path = append(path, s)
	return path
}
