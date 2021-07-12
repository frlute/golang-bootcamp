package mysort

import (
	"sort"

	"github.com/frlute/golang-bootcamp/helper"
)

/*
BubbleSort _
排序算法解释:
	1.比较相邻元素
	2.看是否满足大小关系要求，如果不满足交换位置
	3.循环 1、2 步骤
	4.升序
原地排序, 时间复杂度 O(n^2)
*/
func BubbleSort(data []int) {
	length := len(data)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if data[j] > data[j+1] {
				helper.SwapIntArray(data, j, j+1)
			}
		}
	}
}

// BubbleSortV1 改进：如果数据已排好序，上述实现重重新排序依次，为避免重新排序，对上述算法进行改进，即排序中的自适应特性
func BubbleSortV1(data []int) {
	length := len(data)
	for i := 0; i < length; i++ {
		isChanged := false
		for j := 0; j < length-i-1; j++ {
			if data[j] > data[j+1] {
				helper.SwapIntArray(data, j, j+1)
				isChanged = true
			}
		}
		if !isChanged {
			break
		}
	}
}

/*
BubbleSortV2 _
问题： 上述算法每次交换数据时都是从前像后，但是后面的数据越到后面基本都是已经排好序的，应该可以继续优化
解决方案: 鸡尾酒排序
*/
func BubbleSortV2(data sort.IntSlice) {}
