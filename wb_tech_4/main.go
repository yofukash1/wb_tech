package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func Worker(wg *sync.WaitGroup, ch <-chan int, i int) {
	for a := range ch {
		fmt.Printf("Hello from %d, val: %d\n", i, a)
		time.Sleep(700 * time.Millisecond)
	}
	fmt.Printf("exit%d\n", i)
	wg.Done()
}

func main() {

	var n int
	var wg sync.WaitGroup
	fmt.Scanf("%d", &n)
	done := make(chan struct{})
	channels := make([]chan int, n)

	for i := 0; i < n; i++ {
		ch := make(chan int)
		channels[i] = ch
		wg.Add(1)
		go Worker(&wg, ch, i)

	}

	// to hook Ctrl^c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	wg.Add(1)
	go func(done chan struct{}, wg *sync.WaitGroup) {
		for sig := range c {
			fmt.Println(sig.String())
			done <- struct{}{}
			for _, ch := range channels {
				close(ch)
			}
			wg.Done()
		}
	}(done, &wg)

	// data producing
	i := 0
	for {
		i = (i + 1) % n
		select {
		case channels[i] <- rand.Int() % 27:
		case <-done:
			fmt.Println("done")
			wg.Wait()
			return
		}

	}

}
