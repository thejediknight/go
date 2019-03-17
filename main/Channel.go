package main

import (
	"fmt"
	"math/rand"
	"time"
)

func operation1(ch chan int)  {
	n := rand.Intn(3000)
	time.Sleep(time.Duration(n) * time.Millisecond)
	ch <- n
}

func main() {
	// create a channel for int, atmost holds 1 int.
	ch := make(chan int)

	// Create go routine
	go operation1(ch)
	
	fmt.Println("Waiting for operation1 to complete")

	// Blocks until reads value out of channel.s
	n := <- ch

	fmt.Printf("Process took %dms\n", n)
}