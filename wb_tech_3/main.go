package main

import (
	"fmt"
	"sync"
)

func Adder(ch, out chan int) {
	s := 0
	for n := range ch {
		s += n
	}
	out <- s
}

func main() {
	wg := &sync.WaitGroup{}
	data := []int{2, 4, 6, 8, 10}
	ch := make(chan int)
	out := make(chan int)

	go Adder(ch, out)

	for _, v := range data {
		wg.Add(1)
		go func(val int, out chan<- int) {
			out <- val * val
			wg.Done()
		}(v, ch)
	}
	wg.Wait()
	close(ch)

	s := <-out

	fmt.Print(s)

}
