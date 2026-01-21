package main

import (
	"math"
)

func (n *Node) Reconfigure() {
	primaryLen := len(n.Primary)

	// TODO: If nothing has changed (take a 2 node graph for example) we may want to exponentially backoff the configurations
	// TODO: Maybe each node has a "lazy" factor. When it first joins it's not lazy at all. Eager to reconfigure.
	// 	As long as reconfiguration is happening, lazy maintains. But if we reconfigure and it doesn't result in a change. Make more lazy
	// 	Eventually, it becomes so lazy that no change happens except for rarely. UNTIL another node says "Hey! I'm new!" then lazy is reset?

	// TODO: Before reconfiguration, we really should try and find a space that's semi-optimal to live.

	switch {
	case primaryLen < 2:
		// TODO: Force a reconfiguration
		break
	case primaryLen < 3:
		// TODO: More frequently check if a more optimal configuration is available
		break
	case primaryLen == 3:
		// TODO: Occasionally check if a more optimal configuration is available (Keep it rare. < 3 and < 2 will handle a lot of cases)
		break
	}
}

//func (n *Node) ConfigureV2(start *Node) {
//	if len(start.Primary) < 3 {
//		n.Join(start)
//	} else {
//		n.SpliceBetween(start, start.Primary[0])
//	}
//}

//func (n *Node) Join(other *Node) {
//	n.Primary = append(n.Primary, other)
//	other.Primary = append(other.Primary, n)
//
//	reporter.NewEdge(Edge{
//		From:     n.Id,
//		To:       other.Id,
//		Distance: n.DistanceTo(other),
//	})
//}

//func (n *Node) SpliceBetween(first *Node, second *Node) {
//	fs := first.FindPrimary(second.Id)
//	sf := second.FindPrimary(first.Id)
//
//	reporter.RemoveEdge(Edge{
//		From:     first.Id,
//		To:       second.Id,
//		Distance: 0,
//	})
//
//	n.Primary = append(n.Primary, first)
//	n.Primary = append(n.Primary, second)
//
//	reporter.NewEdge(Edge{
//		From:     first.Id,
//		To:       n.Id,
//		Distance: n.DistanceTo(first),
//	})
//
//	reporter.NewEdge(Edge{
//		From:     second.Id,
//		To:       n.Id,
//		Distance: n.DistanceTo(second),
//	})
//
//	first.Primary[fs] = n
//	second.Primary[sf] = n
//}

// TODO: this doesn't create an optimal configuration at all. It always creates a very large, non cyclical tree.
// 	Nodes should always have 3 primary -- the 3 closest nodes to it

// FindPrimary Returns the index of the primary connection with the specified ID
func (n *Node) FindPrimary(id int) int {
	for i, j := range n.Primary {
		if j.To.Id == id {
			return i
		}
	}

	panic("Somehow tried to find ID that wasn't on a node.")
}

func (n *Node) DistanceTo(node *Node) int {
	xDiff := math.Pow(float64(n.X-node.X), 2)
	yDiff := math.Pow(float64(n.Y-node.Y), 2)

	return int(math.Sqrt(xDiff + yDiff))
}

//func (n *Node) Configure(start *Node) {
//	// If the network only has 1 node!
//	startPrimeLength := len(start.Primary)
//
//	if startPrimeLength == 0 {
//		fmt.Printf("Node %d configured against an empty node %d\n", n.Id, start.Id)
//		start.Primary = append(start.Primary, n)
//		n.Primary = append(n.Primary, start)
//		return
//	}
//
//	dToS := n.DistanceTo(start)
//	minD := NODE_COUNT * NODE_COUNT * 2
//	minI := -1
//
//	// If dToS is less than ALL other values, connect to dToS
//	// Otherwise, call Configure against the closest one
//
//	for i, p := range start.Primary {
//		newD := n.DistanceTo(p)
//		if newD < minD {
//			minD = newD
//			minI = i
//		}
//	}
//
//	//fmt.Printf("Configuring %d against %d. Distance: %d. MinD: %d\n", n.Id, start.Id, dToS, minD)
//
//	// TODO: I don't think this necessarily accounts for local minima. It'll need to also re-organize based on secondaries
//	if dToS <= minD { // "Start" is already the closest node
//		if startPrimeLength < 3 { // We have a non-full primary!
//			start.Primary = append(start.Primary, n)
//			n.Primary = append(n.Primary, start)
//			fmt.Printf("Node %d configured on a semi-empty node %d\n", n.Id, start.Id)
//		} else { // We need to splice the graph
//			other := start.Primary[minI]
//			fmt.Printf("Node %d spliced between %d and %d\n", n.Id, start.Id, other.Id)
//
//			n.Primary = append(n.Primary, other)
//			n.Primary = append(n.Primary, start)
//			other.Primary[other.FindPrimary(start.Id)] = n
//			start.Primary[minI] = n
//		}
//	} else {
//		n.Configure(start.Primary[minI])
//	}
//}
