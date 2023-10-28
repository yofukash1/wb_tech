package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	m  map[int]string
	mu *sync.RWMutex
}

func NewSafeMap() SafeMap {
	return SafeMap{make(map[int]string), &sync.RWMutex{}}
}

func (s *SafeMap) Set(key int, val string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.m[key] = val

}

func main() {
	wg := &sync.WaitGroup{}
	sMap := NewSafeMap()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(sMap *SafeMap, i int) {
			defer wg.Done()
			sMap.Set(i, fmt.Sprint(i))
			fmt.Printf("function %d has set the val %d\n", i, i)
		}(&sMap, i)
	}
	wg.Wait()
	fmt.Print(sMap.m)
}
