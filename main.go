package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mathFunc(-1.0))
}

// y=arcsin(-x) + 13x

func superAsin(x float64) float64 {
	y := math.Asin(-x)
	return y
}

func superMultiple(x float64) float64 {
	return x * 13
}

func mathFunc(x float64) float64 {
	y := superAsin(x) + superMultiple(x)
	return y
}
