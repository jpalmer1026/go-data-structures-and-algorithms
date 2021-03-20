package ds

import "fmt"

// undirected graph
type Graph struct {
	numberOfNodes int
	adjacentList  map[string][]string
}

func NewGraph() *Graph {
	return &Graph{adjacentList: make(map[string][]string)}
}

func (g *Graph) AddVertex(v string) {
	g.adjacentList[v] = []string{}
	g.numberOfNodes++
}

func (g *Graph) AddEdge(v1, v2 string) {
	g.adjacentList[v1] = append(g.adjacentList[v1], v2)
	g.adjacentList[v2] = append(g.adjacentList[v2], v1)
}

func (g *Graph) ShowConnections() {
	fmt.Println(g.adjacentList)
}
