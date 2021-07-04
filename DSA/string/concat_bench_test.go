package string

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkAddStringWithConcat(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < b.N; i++ {
		_ = hello + "," + world
	}
}

func BenchmarkAddStringWithSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s, %s", "hello", "world")
	}
}

func BenchmarkAddStringWithJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Join([]string{"hello", "world"}, ", ")
	}
}

func BenchmarkAddStringWithBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var str bytes.Buffer
		str.WriteString("hello")
		str.WriteString(",")
		str.WriteString("world")
		_ = str.String()
	}
}

func BenchmarkAddStringWithBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var str strings.Builder
		str.WriteString("hello")
		str.WriteString(",")
		str.WriteString("world")
		_ = str.String()
	}
}
