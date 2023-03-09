package overview_test

import (
	"log"
	"testing"
)

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (l *LinkedList) AddToHead(node Node) {
	if node.nextNode == nil {
		node.nextNode = l.headNode
	}
	l.headNode = &node
}

func (l *LinkedList) IterateList() {
	var node *Node
	for node = l.headNode; node != nil; node = node.nextNode {
		log.Println(node.property)
	}
}

func TestLinkedList(t *testing.T) {
	var list LinkedList

	node1 := Node{1, nil}
	node2 := Node{2, nil}
	node3 := Node{3, nil}

	list.AddToHead(node1)
	list.AddToHead(node2)
	list.AddToHead(node3)
	list.IterateList()
}
