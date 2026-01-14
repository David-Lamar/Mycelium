package main

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
	return (a+b)*(a+b+1)/2 + b
}
