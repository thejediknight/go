package main

import (
	"fmt"
)

type DNode struct {
	Value int
	Prev, Next *DNode
}

type DoublyLinkedList struct {
	Head, Tail *DNode
}

func (list *DoublyLinkedList) Add(val int) {

	// Create a new node
	node := &DNode {
		Value: val,
	}

	// If list is empty, point head and to this node
	if list.Head == nil {
		list.Head = node
	} else {
		// Point next of last node to this node, and prev of node to Tail
		list.Tail.Next = node
		node.Prev = list.Tail
	}

	// Finally, point tail to this node
	list.Tail = node		
}

func (list *DoublyLinkedList) Remove(val int) {
	// Closure func to remove curr node from list.
	removeInternal := func (prev *DNode, curr *DNode) {
		// If this is the first node
		if prev == nil {
			// Point head to next node
			list.Head = curr.Next
			
			// If this is the only node
			if curr.Next == nil {
				list.Tail = nil
			} else {
				// Clear prev pointer of next node
				curr.Next.Prev = nil
			}
		} else {
			prev.Next = curr.Next

			// If this is the last node, point Tail to prev
			if curr.Next == nil {
				list.Tail = prev
			} else {
				curr.Next.Prev = prev
			}
		}
	}
	
	curr := list.Head
	var prev *DNode = nil
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

func (list *DoublyLinkedList) PrintList() {
	h := list.Head
	for ;h != nil; {
		fmt.Print(h.Value, " ")
		h = h.Next
	}
	fmt.Println()
}

func (list *DoublyLinkedList) PrintListReverse() {
	t := list.Tail
	for ;t != nil; {
		fmt.Print(t.Value, " ")
		t = t.Prev
	}
	fmt.Println()
}


func main() {

	list := &DoublyLinkedList{}

	list.Add(1)
	list.Add(3)
	list.Add(5)
	list.Add(7)

	list.PrintList()

	list.Remove(3)

	list.PrintList()

	list.PrintListReverse()
}