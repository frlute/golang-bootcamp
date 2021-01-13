package main

import (
	"fmt"
	"math"
)

// Sqrt 主要公式 z -= (z*z-x)/(2*z)
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("cannot Sqrt negativ number: %g", float64(x))
	}
	z := 1.0
	// z := x/2
	const accuracy = 0.0000000000001
	for {
		tmp := z - (z*z-x)/(2*z)
		fmt.Println(tmp)
		if tmp == z || math.Abs(tmp-z) < accuracy {
			break
		}
		z = tmp
	}
	return z, nil
}

func main() {
	fmt.Print(math.Sqrt(5), "=====\n")
	value, err := Sqrt(5)
	if err != nil {
		panic(err)
	}
	fmt.Println(value)
	fmt.Println("pow: ", math.Pow(3, 4))
}
