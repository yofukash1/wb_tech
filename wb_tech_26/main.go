package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)
	sl := make(map[rune]int, len(s))
	for _, v := range s {
		sl[v] += 1
	}

	for _, v := range sl {
		if v > 1 {
			fmt.Print(false)
			return
		}
	}
	fmt.Print(true)

}
