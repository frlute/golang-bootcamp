package mymath

import (
	"fmt"
	"math"
)

// Sqrt 主要公式 z -= (z*z-x)/(2*z)
func Sqrt(x, accuracy float64) (float64, error) {
	if x <= 1 {
		return x, fmt.Errorf("cannot Sqrt negativ number: %g", float64(x))
	}
	z := float64(x / 2)
	for {
		tmp := z - (z*z-x)/(2*z)
		if tmp == z || math.Abs(tmp-z) < 1e-15 {
			break
		}
		z = tmp
	}
	return z, nil
}

// Sqrt2 用二分法实现
func Sqrt2(x int) int {
	if x < 2 {
		return x
	}

	left, right := 1, x/2
	for left <= right {
		mid := left + ((right - left) >> 1)
		if mid == x/mid {
			return mid
		} else if mid < x/mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}
