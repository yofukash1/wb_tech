package main

import "fmt"

func main() {
	elems := []string{"Hello", "Hi", "Bye", "Bye", "1"}
	helper := make(map[string]bool)

	for _, v := range elems {

		if !helper[v] {
			helper[v] = true
		}
	}

	set := make([]string, 0)
	for k := range helper {
		set = append(set, k)
	}
	fmt.Println(set)
}
