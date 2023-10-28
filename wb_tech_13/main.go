package main

import "fmt"

func Swap(a *int, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}

func main() {
	a := 5
	b := 10
	Swap(&a, &b)
	fmt.Printf("%d %d", a, b)
}
