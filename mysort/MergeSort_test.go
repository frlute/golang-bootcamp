package mysort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	cases := []struct {
		data   []int
		expect []int
	}{
		{
			[]int{4, 5, 3, 9, 1},
			[]int{1, 3, 4, 5, 9},
		},
	}

	for _, c := range cases {
		result := MergeSort(c.data)
		if !reflect.DeepEqual(result, c.expect) {
			t.Errorf("MergeSort(%v) expect value: %v, actual value: %v", c.data, c.expect, result)
		}
	}
}
