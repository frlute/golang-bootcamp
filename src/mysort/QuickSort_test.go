package mysort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	cases := []struct {
		data   []int
		expect []int
	}{
		{
			[]int{5, 1, 1, 2, 0, 0},
			[]int{0, 0, 1, 1, 2, 5},
		},
	}

	for _, c := range cases {
		QuickSort(c.data)
		if !reflect.DeepEqual(c.data, c.expect) {
			t.Errorf("QuickSort expect value: %v, actual value: %v", c.expect, c.data)
		}
	}
}
