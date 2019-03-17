package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next *Node
}

type LinkedList struct {
	Head, Tail *Node
}

func (list *LinkedList) Add(val int) {

	// Create a new node
	node := &Node {
		Value: val,
		Next: nil,
	}

	// If list is empty, point head and to this node
	if list.Head == nil {
		list.Head = node
	} else {
		// Point next of last node to this node
		list.Tail.Next = node
	}

	// Finally, point tail to this node
	list.Tail = node		
}

func (list *LinkedList) Remove(val int) {
	// Closure func to remove curr node from list.
	removeInternal := func (prev *Node, curr *Node) {
		// If this is the first node
		if prev == nil {
			list.Head = curr.Next

			// If this is the only node
			if curr.Next == nil {
				list.Tail = nil
			}
		} else {
			prev.Next = curr.Next

			// If this is the last node, point Tail to prev
			if curr.Next == nil {
				list.Tail = prev
			}
		}
	}
	
	curr := list.Head
	var prev *Node = nil
	for ;curr != nil; {
		fmt.Println("Debug: ", curr.Value)
		if curr.Value == val {
			removeInternal(prev, curr)
			break
		} else {
			prev = curr
			curr = curr.Next
		}
	}
}

func (list *LinkedList) PrintList() {
	h := list.Head
	for ;h != nil; {
		fmt.Print(h.Value, " ")
		h = h.Next
	}
	fmt.Println()
}

func main() {

	list := &LinkedList{}

	list.Add(1)
	list.Add(3)
	list.Add(5)
	list.Add(7)

	list.PrintList()

	list.Remove(7)

	list.PrintList()
}