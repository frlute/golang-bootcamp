package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"

	"github.com/frlute/go-playground/cmd/hello/morestrings"
)

func main() {
	fmt.Println("Hello, World.")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello", "Hello"))
}
