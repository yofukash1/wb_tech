package main

import (
	"fmt"
	"math/big"
)

func Input() (a, b *big.Int, err error) {
	a = new(big.Int)
	b = new(big.Int)
	var strA, strB string
	fmt.Scan(&strA)
	fmt.Scan(&strB)

	_, err = fmt.Sscan(strA, a)
	if err != nil {
		return
	}

	_, err = fmt.Sscan(strB, b)
	if err != nil {
		return
	}
	return
}

func Add(a, b *big.Int) (res *big.Int) {
	res = new(big.Int)
	res.Add(a, b)
	return
}

func Sub(a, b *big.Int) (res *big.Int) {
	res = new(big.Int)
	res.Sub(a, b)
	return
}

func Mul(a, b *big.Int) (res *big.Int) {
	res = new(big.Int)
	res.Mul(a, b)
	return
}

func main() {
	a, b, err := Input()
	if err != nil {
		panic(err)
	}

	fmt.Println(Add(a, b))
	fmt.Println(Sub(a, b))
	fmt.Println(Mul(a, b))

}
