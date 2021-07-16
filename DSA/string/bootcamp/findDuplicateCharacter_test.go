package bootcamp

import "testing"

func Test_findDuplicateCharacter(t *testing.T) {
	tests := []struct {
		input string
	}{
		{
			input: "Programming",
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			findDuplicateCharacter(test.input)
		})
	}
}
