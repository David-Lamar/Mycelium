package main

import (
	"fmt"
	lru "github.com/hashicorp/golang-lru/v2"
	"math/rand/v2"
)

type Exec struct {
	fn func(n *Node)
}

type Node struct {
	Id  int
	X   int // X position which determines its "closeness" to other nodes
	Y   int // Y position which determines its "closeness" to other nodes
	Age int

	Primary   []*Node
	Secondary []*Node

	Available bool

	MessageCache *lru.Cache[int, Message]

	RunChan chan Exec
	stopCh  chan struct{}
}

func NewNode(id int) *Node {
	cache, err := lru.New[int, Message](5)
	if err != nil {
		panic(err)
	}

	node := &Node{
		Id:           id,
		X:            rand.IntN(NODE_COUNT * 2),
		Y:            rand.IntN(NODE_COUNT * 2),
		Age:          0,
		Available:    true,
		MessageCache: cache,
		stopCh:       make(chan struct{}),
		RunChan:      make(chan Exec),
	}

	go node.startProcessing()

	return node
}

func (n *Node) startProcessing() {
	defer func() {
		fmt.Printf("Node %d done processing\n", n.Id)
	}()

	fmt.Printf("Node %d starting processing\n", n.Id)

	for {
		select {
		case <-n.stopCh:
			return
		case msg := <-n.RunChan:
			n.exec(msg)
			break
		}
	}

}

func (n *Node) Tick() {
	n.Reconfigure()

	// TODO: Handle misc. things that nodes do while they're "running"
	// 	- Heartbeats, etc.
	// 	- Detect damage
	// 	- Reconfigure
	// 	- etc.

	n.Age++
}

func (n *Node) exec(exec Exec) {
	exec.fn(n)
}

func (n *Node) Stop() {
	close(n.stopCh)
}
