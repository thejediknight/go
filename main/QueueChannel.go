package main

import (
	"fmt"
)

type ChanQueue struct {
	items chan int
}

func (q *ChanQueue) Enqueue(item int) {
	q.items <- item
}

func (q *ChanQueue) Dequeue() (int) {
	return <- q.items
}

func main() {

	// Make the channel with capacity 16
	q := &ChanQueue {
		items: make(chan int, 16),
	}

	q.Enqueue(1)
	q.Enqueue(2)

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}