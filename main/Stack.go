package main

import (
	"fmt"
	"errors"
)

type Stack struct {
	data []int
}

func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() (int, error) {
	dataLength := len(s.data)

	if (dataLength == 0) {
		return -1, errors.New("Err: Stack is empty!")
	}

	// Multi assignment
	poppedValue, newStackData := s.data[dataLength-1], s.data[0:dataLength-1]
	
	s.data = newStackData
	return poppedValue, nil
}

func main() {
	stack := &Stack{}

	stack.Push(1)
	stack.Push(2)

	if val, err := stack.Pop(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}
}
