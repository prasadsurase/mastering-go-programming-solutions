// singly linked list

package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func (n *Node) GetValue() int {
	return n.value
}

func (n *Node) SetValue(val int) int {
	n.value = val
	return val
}

func NewNode() *Node {
	return new(Node)
}

func main() {
	start := NewNode()
	start.SetValue(1)

	next := start

	for i := 2; i <= 10; i++ {
		node := NewNode()
		node.SetValue(i)
		next.next = node
		next = node
	}

	node := start
	for {
		fmt.Printf("%v\n", node)
		if node.next == nil {
			break
		} else {
			node = node.next
		}
	}
}
