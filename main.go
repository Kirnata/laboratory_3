package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello! It's mini calculator for y=arcsin(-x) + 13x")
	x := -2.0
	for x < -1.0 || x > 1.0 {
		fmt.Println("Print some in [-1;1]")
		if _, err := fmt.Scan(&x); err != nil {
			fmt.Errorf("%s", err)
			return
		}
	}
	fmt.Println(mathFunc(x))
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
