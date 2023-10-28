package main

import (
	"fmt"
	"time"
)

func main() {
	data := []int{2, 4, 6, 8, 10}

	for _, v := range data {
		go func(val int) {
			fmt.Printf("%d\n", val*val)
		}(v)
	}
	time.Sleep(time.Second * 2)
}
