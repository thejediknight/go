package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleep_fo() {
	n := rand.Intn(3000)
	time.Sleep(time.Duration(n) * time.Millisecond)
}

//~ ch is write-only channel
func producer_fo(ch chan <- int) {
	for {
		// sleep for some random time
		sleep_fo()

		// Generate some random number
		n := rand.Intn(100)

		fmt.Printf("Sending %d\n", n)
		ch <- n
	}
}

//~ ch is read-only channel
func consumer_fo(ch <- chan int, name string) {
	for n:= range ch {
		fmt.Printf("Consumer %s <- %d\n", name, n)
	}
}

//# Read from channel chA and write to chC and chB 
func fanOut(chA <- chan int, chB, chC chan <- int)  {
	for n:= range chA {
		if n < 50 {
			chB <- n
		} else {
			chC <- n
		}
	}
}

func FanOutOperations() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	go producer_fo(chA)
	go consumer_fo(chB, "B")
	go consumer_fo(chC, "c")

	fanOut(chA, chB, chC)
}