package main

import (
	"fmt"
)

func fibonacci() func() int {
	var a, b int = -1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
