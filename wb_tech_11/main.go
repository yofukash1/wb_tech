package main

import (
	"fmt"
	"math/rand"
)

func Intersec(set1, set2 map[int]bool, dest map[int]bool) {
	if len(set1) <= len(set2) {
		for k := range set1 {
			if set1[k] == set2[k] {
				dest[k] = true
			}
		}
	} else {
		for k := range set2 {
			if set1[k] == set2[k] {
				dest[k] = true
			}
		}
	}
}

func Init(set1, set2 map[int]bool) {
	for i := 0; i < 23; i++ {
		key := rand.Int() % 27
		set1[key] = true

	}

	for i := 0; i < 19; i++ {
		key := rand.Int() % 27
		set2[key] = true
	}
}

func main() {
	set1 := make(map[int]bool, 13)
	set2 := make(map[int]bool, 9)
	intersec := make(map[int]bool)
	Init(set1, set2)

	Intersec(set1, set2, intersec)

	fmt.Println(set1)
	fmt.Println(set2)
	fmt.Println(intersec)

}
