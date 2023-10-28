package main

import (
	"fmt"
	"math/rand"
	"time"
)

func DataProducer(out chan<- int, done chan struct{}) {
	for {
		value := rand.Int() % 27
		time.Sleep(500 * time.Millisecond)
		select {
		case out <- value:
			fmt.Printf("Produced val: %d \n", value)
		case <-done:
			fmt.Println("Producing edned")
		}
	}
}

func Worker(in <-chan int, out chan<- int) {
	for val := range in {
		val += 1
		out <- val
	}
}

func main() {
	in := make(chan int)
	out := make(chan int)
	done := make(chan struct{})

	var n int
	fmt.Scanf("%d", &n)

	go DataProducer(out, done)
	go Worker(out, in)

	timeout := time.After(time.Duration(n) * time.Second)
	for {
		select {
		case val := <-in:
			fmt.Printf("processed val: %d\n", val)
		case <-timeout: // strictly does not work??
			done <- struct{}{}
			fmt.Println("time is out")
			close(in)
			close(out)
			return
		}
	}
}
