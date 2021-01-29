package binarySearch

// FirstGTEBSearch first greater than or equal 在有序数组中，查找第一个大于等于给定值的元素
func FirstGTEBSearch(items []int, target int) int {
	length := len(items)
	if length == 0 {
		return -1
	}

	left, right := 0, length-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		// 更优版使用位运算，待理解 mid := low + ((higt-low)>>1)
		value := items[mid]
		if value >= target {
			if mid == 0 || items[mid-1] < value {
				return mid
			}
			right = mid - 1
		} else if value < target {
			left = mid + 1
		}
	}

	return -1
}
