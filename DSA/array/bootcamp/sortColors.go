package bootcamp

import "github.com/frlute/golang-bootcamp/helper"

/*
题目: [75] 荷兰国旗问题 dutch flag partition
详见: https://leetcode.com/problems/sort-colors/
题解：https://cloud.tencent.com/developer/article/1624933
解题思路:
	- 可以把数据分三块区域，红白蓝，也就是需要两根线
	- 然后提供一个基准移动位置
*/

func sortColors(items []int) []int {
	length := len(items)
	if length <= 1 {
		return items
	}

	left, right := 0, length-1

	for pivot := 0; pivot <= right; pivot++ {
		value := items[pivot]
		// 三个颜色区间
		if value == 0 {
			helper.SwapIntArray(items, left, pivot)
			left++
		} else if value == 2 {
			helper.SwapIntArray(items, pivot, right)
			// 切换完当前元素需要重新判断
			pivot--
			right--
		}
	}

	return items
}
