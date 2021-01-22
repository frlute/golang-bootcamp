package mysort

import (
	"reflect"
	"testing"
)

func TestBucketSort(t *testing.T) {
	cases := []struct {
		data   []int
		expect []int
	}{
		{
			[]int{31, 16, 37, 2, 13, 32, 10, 27, 7, 42, 29, 18, 28, 12, 9},
			[]int{2, 7, 9, 10, 12, 13, 16, 18, 27, 28, 29, 31, 32, 37, 42},
		},
	}

	for _, c := range cases {
		BucketSort(c.data)
		if !reflect.DeepEqual(c.expect, c.data) {
			t.Errorf("BucketSort expect value: %v, actual value: %v", c.expect, c.data)
		}
	}
}
