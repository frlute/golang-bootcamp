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
			[]int{7, 2, 18, 42, 32, 31, 27, 16, 9, 37, 10, 29, 12, 28, 13},
			[]int{2, 7, 9, 10, 12, 13, 16, 18, 27, 28, 29, 31, 32, 37, 42},
		},
	}

	for _, c := range cases {
		BucketSort(c.data)
		if !reflect.DeepEqual(c.expect, c.expect) {
			t.Errorf("BucketSort expect value: %v, actual value: %v", c.expect, c.data)
		}
	}
}
