package mysort

import "sort"

/*
排序算法解释:

*/
func BubbleSort(data sort.Interface) {
	n := data.Len()
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if !data.Less(j, j+1) {
				data.Swap(j, j+1)
			}
		}
	}
}

/*
BubbleSortV1

*/
func BubbleSortV1(data sort.Interface) {
	n := data.Len()
	if n <= 1 {
		return
	}

	isChanged := false
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if !data.Less(j, j+1) {
				data.Swap(j, j+1)
				isChanged = true
			}
		}
		if !isChanged {
			break
		}
	}
}
