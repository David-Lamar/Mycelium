package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

// Handles up to 1 million fairly well during configure

const NODE_COUNT = 100

func main() {

	// TODO: Each node needs to have a way to communicate to its neighbors
	// TODO: When each node is added, it should be added to a single node and move through the network to its optimal spot
	// TODO: When a node is added, it's _provided_ to a specific node but not connected. It will move through the network to find optimal placement
	// TODO: Handle deleting a node and how a node will handle that error
	// TODO: Handle communication between nodes (heartbeat kind of thing)

	var nodes []*Node

	for i := 0; i <= NODE_COUNT; i++ {
		node := &Node{
			Id: i,
			X:  rand.IntN(NODE_COUNT * 10),
			Y:  rand.IntN(NODE_COUNT * 10),
		}

		nodes = append(nodes, node)

		if i > 0 {
			node.Configure(nodes[0])
		}
	}
}

type Node struct {
	Id int
	X  int // X position which determines its "closeness" to other nodes
	Y  int // Y position which determines its "closeness" to other nodes

	Primary   []*Node
	Secondary []*Node
}

func (n *Node) DistanceTo(node *Node) int {
	xDiff := math.Pow(float64(n.X-node.X), 2)
	yDiff := math.Pow(float64(n.Y-node.Y), 2)

	return int(math.Sqrt(xDiff + yDiff))
}

func (n *Node) Configure(start *Node) {
	// If the network only has 1 node!
	startPrimeLength := len(start.Primary)

	if startPrimeLength == 0 {
		fmt.Printf("Node %d configured against an empty node %d\n", n.Id, start.Id)
		start.Primary = append(start.Primary, n)
		n.Primary = append(n.Primary, start)
		return
	}

	dToS := n.DistanceTo(start)
	minD := NODE_COUNT * 10 * 3
	minI := -1

	// If dToS is less than ALL other values, connect to dToS
	// Otherwise, call Configure against the closest one

	for i, p := range start.Primary {
		newD := n.DistanceTo(p)
		if newD < minD {
			minD = newD
			minI = i
		}
	}

	//fmt.Printf("Configuring %d against %d. Distance: %d. MinD: %d\n", n.Id, start.Id, dToS, minD)

	// TODO: I don't think this necessarily accounts for local minima. It'll need to also re-organize based on secondaries
	if dToS <= minD { // "Start" is already the closest node
		if startPrimeLength < 3 { // We have a non-full primary!
			start.Primary = append(start.Primary, n)
			n.Primary = append(n.Primary, start)
			fmt.Printf("Node %d configured on a semi-empty node %d\n", n.Id, start.Id)
		} else { // We need to splice the graph
			other := start.Primary[minI]
			fmt.Printf("Node %d spliced between %d and %d\n", n.Id, start.Id, other.Id)

			n.Primary = append(n.Primary, other)
			n.Primary = append(n.Primary, start)
			other.Primary[other.FindPrimary(start.Id)] = n
			start.Primary[minI] = n
		}
	} else {
		n.Configure(start.Primary[minI])
	}
}

// FindPrimary Returns the index of the primary connection with the specified ID
func (n *Node) FindPrimary(id int) int {
	for i, j := range n.Primary {
		if j.Id == id {
			fmt.Printf("Found primary at position %d\n", i)
			return i
		}
	}

	panic("Somehow tried to find ID that wasn't on a node.")
}
