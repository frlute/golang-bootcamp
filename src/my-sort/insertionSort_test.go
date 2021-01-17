package mysort

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	cases := []struct {
		data   []int
		expect []int
	}{
		{
			[]int{4, 5, 3, 9, 1},
			[]int{1, 34, 5, 9},
		},
	}

	for _, c := range cases {
		InsertionSort(c.data)
		if reflect.DeepEqual(c.data, c.expect) {
			t.Errorf("InsertionSort expect value: %v, expect value: %v", c.expect, c.data)
		}
	}
}
