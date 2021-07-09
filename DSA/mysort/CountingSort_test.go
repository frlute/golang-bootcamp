package mysort

import (
	"reflect"
	"testing"
)

func TestCountingSort(t *testing.T) {
	cases := []struct {
		data   []int
		expect []int
	}{
		{
			[]int{11, 9, 6, 8, 1, 3, 5, 1, 1, 0, 100},
			[]int{0, 1, 1, 1, 3, 5, 6, 8, 9, 11, 100},
		},
	}

	for _, c := range cases {
		result := CountingSort(c.data)
		if !reflect.DeepEqual(c.expect, result) {
			t.Errorf("CountingSort(%v) expect value: %v, actual value: %v", c.data, c.expect, result)
		}
	}

	for _, c := range cases {
		CountingSort2(c.data)
		if !reflect.DeepEqual(c.expect, c.data) {
			t.Errorf("CountingSort2 expect value: %v, actual value: %v", c.expect, c.data)
		}
	}
}
