package graphs

import "fmt"

type Graph struct {
	vertices  int32
	edges     int32
	adjacency map[string]map[string]int32
}

func (g *Graph) Print() {
	fmt.Println(g.adjacency)
}

func (g *Graph) Vertices() []string {
	verticesList := make([]string, 0)
	if g.adjacency != nil {
		for key, _ := range g.adjacency {
			verticesList = append(verticesList, key)
		}
	}
	return verticesList
}

func (g *Graph) Adjacency(v string) []string {
	adjacencyList := make([]string, 0)
	if g.adjacency != nil {
		if list, exist := g.adjacency[v]; exist {
			for adj, _ := range list {
				adjacencyList = append(adjacencyList, adj)
			}
		}
	}
	return adjacencyList
}

func (g *Graph) AddEdge(v1, v2 string) {
	if g.adjacency == nil {
		g.adjacency = make(map[string]map[string]int32)
	}
	if _, exist := g.adjacency[v1]; !exist {
		g.adjacency[v1] = make(map[string]int32)
	}
	g.adjacency[v1][v2] = 1
	if _, exist := g.adjacency[v2]; !exist {
		g.adjacency[v2] = make(map[string]int32)
	}
	g.adjacency[v2][v1] = 1
}

func (g *Graph) AddEdgeDistance(v1, v2 string, distance int32) {
	if g.adjacency == nil {
		g.adjacency = make(map[string]map[string]int32)
	}

	if _, exist := g.adjacency[v1]; !exist {
		g.adjacency[v1] = make(map[string]int32)
	}
	g.adjacency[v1][v2] = distance
	if _, exist := g.adjacency[v2]; !exist {
		g.adjacency[v2] = make(map[string]int32)
	}
	g.adjacency[v2][v1] = distance
}

func (g *Graph) AddEdgeUnidirectDistance(v1, v2 string, distance int32) {
	if g.adjacency == nil {
		g.adjacency = make(map[string]map[string]int32)
	}

	if _, exist := g.adjacency[v1]; !exist {
		g.adjacency[v1] = make(map[string]int32)
	}
	g.adjacency[v1][v2] = distance
}

func (g *Graph) RemoveEdge(v1, v2 string) {
	if _, exist := g.adjacency[v1]; exist {
		if _, exist := g.adjacency[v1][v2]; exist {
			delete(g.adjacency[v1], v2)
		}
	}
	if _, exist := g.adjacency[v2]; exist {
		if _, exist := g.adjacency[v2][v1]; exist {
			delete(g.adjacency[v2], v1)
		}
	}

}
