package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isIsomorphic(t *testing.T) {

	cases := []struct {
		source string
		target string
		output bool
	}{
		{
			source: "egg",
			target: "add",
			output: true,
		},
		{
			source: "foo",
			target: "bar",
			output: false,
		},
		{
			source: "paper",
			target: "title",
			output: true,
		},
	}

	for _, item := range cases {
		t.Run("", func(t *testing.T) {
			// res := isIsomorphic(item.source, item.target)
			// assert.Equal(t, item.output, res)

			assert.Equal(t, item.output, isIsomorphicBasedHash(item.source, item.target))
		})
	}
}
