package main

import (
	"fmt"
	"math/rand"
	"time"

	"sort"
)

func main() {
	s := make([]int, 32)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		s[i] = rand.Intn(32)
	}
	sort.Ints(s)
	a := rand.Intn(32)
	fmt.Printf("Sorted list:%v, val: %d\n", s, a)
	pos, in := BinSearch(s, a)
	if in {
		fmt.Printf("Position: %d", pos)
	} else {
		fmt.Println("value is not in the array")
	}
}

func BinSearch(list []int, item int) (position int, in bool) {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]
		if guess == item {
			return mid, true
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0, false
}
