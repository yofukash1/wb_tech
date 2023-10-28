package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func Worker(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	for n := range in {
		out <- n * n
	}
	close(out)
	wg.Done()
}

func Output(in <-chan int, wg *sync.WaitGroup) {
	for output := range in {
		fmt.Println(output)
	}
	wg.Done()
}

func Init(nums []int) {
	for i := 0; i < 49; i++ {
		nums[i] = rand.Int() % 49
	}
}

func main() {
	wg := sync.WaitGroup{}
	nums := make([]int, 49)
	Init(nums)
	in := make(chan int)
	out := make(chan int)

	wg.Add(2)
	go Worker(in, out, &wg)
	go Output(out, &wg)
	for v := range nums {
		in <- v
	}
	close(in)

	wg.Wait()
}
