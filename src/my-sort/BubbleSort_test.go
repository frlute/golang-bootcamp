package mysort

import (
	"reflect"
	"sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	cases := []struct {
		data   sort.IntSlice
		expect sort.IntSlice
	}{
		{
			sort.IntSlice{22, 34, 3, 40, 18, 4},
			sort.IntSlice{3, 4, 18, 22, 34, 40},
		},
	}

	for _, v := range cases {
		BubbleSort(v.data)
		if !reflect.DeepEqual(v.data, v.expect) {
			t.Errorf("BubbleSort expect value: %v, actual value: %+v", v.expect, v.data)
		}
	}

	for _, v := range cases {
		BubbleSortV1(v.data)
		if !reflect.DeepEqual(v.data, v.expect) {
			t.Errorf("BubbleSortV1 expect value: %v, actual value: %+v", v.expect, v.data)
		}
	}
}
