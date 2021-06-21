package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"

	"github.com/frlute/golang-bootcamp/cmd/hello/morestrings"
)

func main() {
	fmt.Println("Hello, World.")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello", "Hello"))
}
