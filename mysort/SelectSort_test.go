package mysort

import (
	"reflect"
	"testing"
)

func TestSelectSort(t *testing.T) {
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
		SelectSort(c.data)
		if !reflect.DeepEqual(c.data, c.expect) {
			t.Errorf("SelectSort expect value: %v, expect value: %v", c.expect, c.data)
		}
	}
}
