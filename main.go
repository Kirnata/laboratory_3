package main

import (
	"math"
)

func main() {

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
