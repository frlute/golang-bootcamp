package main

import (
	"fmt"

	s "github.com/frlute/go-playground/cmd/search/binarySearch"
)

func main() {
	data := []int{1, 3, 5, 7, 9}
	target := 3
	fmt.Println(s.Search(data, target))
}
