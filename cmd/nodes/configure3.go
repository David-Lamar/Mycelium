package main

func (n *Node) ConfigureV3(start *Node) {
	dToS := n.DistanceTo(start)

	minD := dToS
	minI := -1

	for i, p := range start.Primary {
		newD := n.DistanceTo(p.To)
		if newD < minD {
			minD = newD
			minI = i
		}
	}

	if dToS <= minD { // "Start" is already the closest node
		n.ConnectTo(start)
	} else {
		n.ConfigureV3(start.Primary[minI].To)
	}
}

func (n *Node) CreateConnectionBetween(other *Node) {
	dist := n.DistanceTo(other)

	conn1 := Connection{
		To:     other,
		Weight: dist,
	}

	conn2 := Connection{
		To:     n,
		Weight: dist,
	}

	other.Primary = append(other.Primary, conn2)
	n.Primary = append(n.Primary, conn1)
}

func (n *Node) ConnectTo(target *Node) {
	otherPrimeLen := len(target.Primary)

	if otherPrimeLen == 0 {
		// This is the only other node in the network
		n.CreateConnectionBetween(target)
		return
	}

	// TODO:
	// 	1. Get all of the nodes (that we care about -- ONLY target and its primaries)
	// 	2. Get all of the (existing) connections between nodes
	// 	3.

}
