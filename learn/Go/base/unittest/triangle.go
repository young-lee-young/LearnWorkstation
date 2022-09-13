package main

import (
	"math"
)

func main() {
	
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	return c
}
