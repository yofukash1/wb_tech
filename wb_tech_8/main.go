package main

import (
	"fmt"
)

func InvertInto1(n *int64, i int) {
	*n |= 1 << (i - 1)
	return
}

func InvertInto0(n *int64, i int) {
	*n &^= 1 << (i - 1)
	return
}

func main() {
	var n int64 = 121
	fmt.Printf("%010b\n", n)
	InvertInto1(&n, 9)
	fmt.Printf("%010b\n", n)

	var l int64 = 1024 - 1
	fmt.Printf("%010b\n", l)
	InvertInto0(&l, 9)
	fmt.Printf("%010b\n", l)

}
