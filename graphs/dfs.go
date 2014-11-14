package graphs

type DepthFirstSearch struct {
	marked map[string]bool
}

func (dfs *DepthFirstSearch) Init(g *Graph) {
	if dfs.marked == nil {
		dfs.marked = make(map[string]bool)
	}
	for _, v := range g.Vertices() {
		dfs.marked[v] = false
	}

}

func (dfs *DepthFirstSearch) DFS(g *Graph, v string) {
	dfs.marked[v] = true
	for _, av := range g.Adjacency(v) {
		if !dfs.marked[av] {
			dfs.DFS(g, av)
		}
	}
}

func (dfs *DepthFirstSearch) isConnected() bool {
	for _, val := range dfs.marked {
		if val == false {
			return false
		}
	}
	return true
}

func IsConnected(g *Graph, v string) bool {
	dfsContext := new(DepthFirstSearch)
	dfsContext.Init(g)
	dfsContext.DFS(g, v)
	return dfsContext.isConnected()
}
