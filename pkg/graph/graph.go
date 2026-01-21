package graph

import (
	"Mycelium/pkg/reporter"
	"github.com/gin-gonic/gin"
	"slices"
	"sort"
)

// This is for a node's internal representation of a graph
// Since each node will be stand-alone and won't have access to the real graph, it needs to build an internal representation

type Graph struct {
	nodes map[int]*Node
	edges map[int]*Edge
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[int]*Node),
		edges: make(map[int]*Edge),
	}
}

type Node struct {
	Id    int
	Edges map[int]*Edge
}

type Edge struct {
	id            int
	first, second *Node
	weight        int
}

func (g *Graph) Export() gin.H {
	nodes := slices.Collect(func(yield func(int) bool) {
		for _, n := range g.nodes {
			if !yield(n.Id) {
				return
			}
		}
	})

	edges := slices.Collect(func(yield func(edge reporter.Edge) bool) {
		for _, e := range g.edges {
			edge := reporter.Edge{
				Id:       e.id,
				From:     e.first.Id,
				To:       e.second.Id,
				Distance: e.weight,
			}

			if !yield(edge) {
				return
			}
		}
	})

	return gin.H{
		"Nodes": nodes,
		"Edges": edges,
	}
}

func (g *Graph) FullyConnected(distance func(int, int) int, nodes ...int) {
	for _, n := range nodes {
		g.AddNode(n)
	}

	// Creates an edge between every node
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			g.AddEdge(nodes[i], nodes[j], distance(nodes[i], nodes[j]))
		}
	}
}

func (g *Graph) PartiallyConnected(distance func(int, int) int, newId int, existingId int) {
	target := g.nodes[existingId]

	var nodesToAdd []*Node
	nodesToAdd = append(nodesToAdd, target)

	g.AddNode(newId)

	for _, e := range target.Edges {
		if e.first == target {
			nodesToAdd = append(nodesToAdd, e.second)
		} else {
			nodesToAdd = append(nodesToAdd, e.first)
		}
	}

	for _, n := range nodesToAdd {
		g.AddEdge(newId, n.Id, distance(newId, n.Id))
	}
}

func (g *Graph) AddNode(id int) {
	g.nodes[id] = &Node{
		Id:    id,
		Edges: make(map[int]*Edge),
	}
}

func (g *Graph) AddEdge(node1 int, node2 int, weight int) {
	n1, ok := g.nodes[node1]
	if !ok {
		panic("Cannot add an edge to a non-existent node")
	}

	n2, ok := g.nodes[node2]
	if !ok {
		panic("Cannot add an edge to a non-existent node")
	}

	id := getId(node1, node2)
	edge := &Edge{
		id:     id,
		first:  n1,
		second: n2,
		weight: weight,
	}

	g.edges[id] = edge
	n1.Edges[id] = edge
	n2.Edges[id] = edge
}

func (g *Graph) RemoveEdge(node1 int, node2 int) {
	n1, ok := g.nodes[node1]
	if !ok {
		panic("Cannot remove an edge from a non-existent node")
	}

	n2, ok := g.nodes[node2]
	if !ok {
		panic("Cannot remove an edge from a non-existent node")
	}

	id := getId(node1, node2)

	delete(n1.Edges, id)
	delete(n2.Edges, id)
	delete(g.edges, id)
}

func (g *Graph) Disconnect() {
	var overConnected []*Node
	var edges = make(map[int]*Edge)

	// Step 1. Get all nodes with greater than 3 connections
	for _, n := range g.nodes {
		if len(n.Edges) > 3 {
			overConnected = append(overConnected, n)

			for _, e := range n.Edges {
				edges[e.id] = e
			}
		}
	}

	if len(overConnected) < 1 {
		return
	}

	// Step 2. Get edges of those nodes sorted in ascending weight order
	var edgeList []*Edge

	for _, e := range edges {
		edgeList = append(edgeList, e)
	}

	// Sort in ascending order by weight
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i].weight > edgeList[j].weight
	})

	// Step 3. Remove edges
	removed := false
	for _, e := range edgeList {
		n1, n2 := e.first, e.second
		if len(n1.Edges) > 3 && len(n2.Edges) > 3 {
			g.RemoveEdge(n1.Id, n2.Id)
			removed = true
			break
		}
	}

	if !removed {
		for _, e := range edgeList {
			n1, n2 := e.first, e.second
			if len(n1.Edges) > 2 && len(n2.Edges) > 2 {
				g.RemoveEdge(n1.Id, n2.Id)
				break
			}
		}
	}
}

func getId(a int, b int) int {
	if a < b {
		return (a+b)*(a+b+1)/2 + b
	} else {
		return (b+a)*(b+a+1)/2 + a
	}
}

// TODO: **********************************************************************

// TODO: Automatically sort the nodes by most connections to least
// TODO: Then when a connection is removed, adjust the sorting
// TODO: This makes the query for which node(s) are most connected constant time
// TODO: Sort edges _ON_ a node from worst to best. This way we only ever have to look at X edges for X nodes to determine the worst one.
// TODO: Then, our calculations become very efficient.

// TODO: **********************************************************************

// The Flow:
// 1. Identify how many connections we need to remove
// 2. Get all edges for all nodes that have more than 4 edges
// 3. Get the worst edge from that list of nodes
// 4. If removing the edge doesn't result in a node having less than 2 connections, remove it
// 5. Repeat from 2
