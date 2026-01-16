package main

import (
	"fmt"
	"time"
)

type MessageType byte

const (
	Hello MessageType = 0x00
)

type Message struct {
	Id         int // Message ID
	Originator int // Who this message is _originally_ from
	From       int // Who this message is from (last hop)
	Hop        int
	Type       MessageType // The type of message this is
}

// TODO: This should happen during a tick. Otherwise nodes will work more quickly than we can simulate

func (n *Node) SendMessage(t MessageType) {
	message := Message{
		Id:         int(time.Now().UnixMilli()),
		Originator: n.Id,
		From:       n.Id,
		Type:       t,
		Hop:        0,
	}

	for _, node := range n.Primary {
		node.ReceiveMessage(message)
	}
}

func (n *Node) PropagateMessage(message Message) {
	dontSend := message.From

	message.From = n.Id
	message.Hop = message.Hop + 1

	for _, node := range n.Primary {
		if node.Id == dontSend {
			continue
		}

		node.ReceiveMessage(message)
	}
}

func (n *Node) ReceiveMessage(message Message) {
	handled := n.MessageCache.Contains(message.Id)
	fmt.Printf("Node %d received message %d. Handled? %t\n", n.Id, message.Id, handled)
	if handled {
		// We've already handled this message. No need to again
		return
	}

	n.MessageCache.Add(message.Id, message)

	switch message.Type {
	case Hello:
		n.HandleHello(message)
		break
	}
}

func (n *Node) HandleHello(message Message) {
	n.PropagateMessage(message)

	// TODO: _ALWAYS_ handle it on hop 0. Otherwise, randomly decrease the further out it is and take a sampling
}
