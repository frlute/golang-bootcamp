package mysort

import "sort"

/*
排序算法解释:
	1.比较相邻元素
	2.看是否满足大小关系要求，如果不满足交换位置
	3.循环 1、2 步骤
*/
func BubbleSort(data sort.IntSlice) {
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

// BubbleSortV1 改进：如果某次没有交换时表示已排好序
func BubbleSortV1(data sort.IntSlice) {
	n := data.Len()
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		isChanged := false
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
