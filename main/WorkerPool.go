package main

import (
	"fmt"
	"time"
	"math/rand"
)

func echoWorker(in <- chan int, out chan <- int)  {
	for {
		// Read from queue
		n := <- in

		// Do some processing
		time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

		out <- n	
	}
}

// write-only channel
func produce(ch chan <- int) {
	n := 0
	for {
		fmt.Printf("-> Sending job %d to channel\n", n)
		ch <- n
		n++;
	}
}

func main() {
	in := make(chan int)
	out := make(chan int)

	// create 4 workers 
	for index := 0; index < 4; index++ {
		go echoWorker(in, out)		
	}

	go produce(in)

	for n:= range out {
		fmt.Printf("<- Recv job %d\n", n)
	}
}