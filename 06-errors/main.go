// Video solution: https://www.youtube.com/watch?v=m1O8LtW-LuQ

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number:%v\n", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for {
		prev := z
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-prev) < 10e-4 {
			return z, nil
		}
	}
}

func main() {
	s, err := Sqrt(2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("sqrt:", s)
}
