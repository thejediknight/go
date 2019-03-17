package main

import (
	"fmt"
	"errors"
)

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (int, error) {
	itemsLen := len(q.items)

	if itemsLen == 0 {
		return -1, errors.New("Err: Queue is empty!")
	}

	returnVal, newItems := q.items[0], q.items[1:]
	q.items = newItems

	return returnVal, nil
}

func QueueOperations() {
	q := &Queue{}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	for i:=5; i>=0; i-- {
		if item, err := q.Dequeue(); err == nil {
			fmt.Println(item)
		} else {
			fmt.Println(err)
		}
	}
}