package main

import (
	"fmt"
	"sync"
)

type SafeInc struct {
	i  int
	mu *sync.RWMutex
}

func NewSafeInc() SafeInc {
	return SafeInc{0, &sync.RWMutex{}}
}

func (s *SafeInc) Inc() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.i += 1
}

func main() {
	wg := &sync.WaitGroup{}
	sInc := NewSafeInc()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(sInc *SafeInc, i int) {
			defer wg.Done()
			sInc.Inc()
			fmt.Printf("function %d has incremented\n", i)
		}(&sInc, i)
	}
	wg.Wait()
	fmt.Printf("%d", sInc.i)
}
