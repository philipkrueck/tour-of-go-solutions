// Video explanation: https://www.youtube.com/watch?v=taRekUoZD8I

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	const tolerance = 0.0001
	z := 1.0

	for i := 0; i < 10; i++ {
		prev := z
		z -= (z*z - x) / (2 * z)
		fmt.Println("i =", i, "; z =", z)

		if diff := math.Abs(prev - z); diff < tolerance {
			break
		}
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
