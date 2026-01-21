package reporter

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Edge struct {
	Id       int
	From     int
	To       int
	Distance int
}

type Reporter struct {
	nodes []int
	edges []Edge

	// TODO: Have "events" that span 1 "tick"
	// 	Each pull from the web side will retrieve all of the events that happened in the previous tick
	// 	Then that list will get cleared to get ready for the next one. This will batch things like communication between nodes, etc.
}

func NewReporter() *Reporter {
	return &Reporter{
		nodes: make([]int, 0),
		edges: make([]Edge, 0),
	}
}

func (r *Reporter) NewNode(id int) {
	r.nodes = append(r.nodes, id)
}

func (r *Reporter) NewEdge(edge Edge) {
	edge.Id = getId(edge.To, edge.From)
	fmt.Printf("Edge ID %d\n", edge.Id)
	r.edges = append(r.edges, edge)
}

func (r *Reporter) RemoveEdge(edge Edge) {
	for i, e := range r.edges {
		if (e.To == edge.To && e.From == edge.From) || (e.To == edge.From && e.From == edge.To) {
			println(len(r.edges))
			r.edges = append(r.edges[:i], r.edges[i+1:]...)
			println("REmoved edge!")
			println(len(r.edges))
			return
		}
	}
}

func (r *Reporter) GetDoc() gin.H {
	return gin.H{
		"Nodes": r.nodes,
		"Edges": r.edges,
	}
}

func getId(a int, b int) int {
	if a < b {
		return (a+b)*(a+b+1)/2 + b
	} else {
		return (b+a)*(b+a+1)/2 + a
	}
}
