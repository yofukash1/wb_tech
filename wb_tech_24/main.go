package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

func D(p1, p2 *Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(float64(dx*dx + dy*dy))
}

func main() {
	p1 := NewPoint(1, 1)
	p2 := NewPoint(4, 5)
	fmt.Printf("%v", D(p1, p2))
}
