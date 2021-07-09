package string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isAlphabet(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		input  byte
		output bool
	}{
		{
			input:  'a',
			output: true,
		},
		{
			input:  '&',
			output: false,
		},
		{
			input:  'B',
			output: true,
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("test isAlphabet(%q)", test.input)
		t.Run(name, func(t *testing.T) {
			as.Equal(test.output, isAlphabet(test.input))
		})
	}
}

func Test_reverseStingWithoutAffectSpecialChar(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "a,b$c",
			output: "c,b$a",
		},
		{
			input:  "Ab,c,de!$",
			output: "ed,c,bA!$",
		},
		{
			input:  "a!!!b.c.d,e'f,ghi",
			output: "i!!!h.g.f,e'd,cba",
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("test reverseStingWithoutAffectSpecialChar(%s)", test.input)
		t.Run(name, func(t *testing.T) {
			res := reverseStringWithoutAffectSpecialChar(test.input)
			as.Equal(test.output, res)
		})
	}
}
