package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleep() {
	n := rand.Intn(3000)
	time.Sleep(time.Duration(n) * time.Millisecond)
}

//~ ch is write-only channel
func producer(ch chan <- int, name string) {
	for {
		// sleep for some random time
		sleep()

		// Generate some random number
		n := rand.Intn(100)

		fmt.Printf("Producer %s sending %d\n", name, n)
		ch <- n
	}
}

//~ ch is read-only channel
func consumer(ch <- chan int) {
	for n:= range ch {
		fmt.Printf("<- %d\n", n)
	}
}

//# Read from channels chA and chB, and write to chC
func fanIn(chA, chB <- chan int, chC chan <- int)  {
	var n int

	for {
		select{
		case n = <- chA:
			chC <- n
		case n = <- chB:
			chC <- n
		}
	}	
}

func FanInOperations() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	go producer(chA, "A")
	go producer(chB, "B")
	go consumer(chC)

	fanIn(chA, chB, chC)
}