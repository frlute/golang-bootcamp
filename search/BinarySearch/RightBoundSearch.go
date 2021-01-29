package binarySearch

// RightBoundSearch 表示从右边界开始搜索, such as [1,2,2,2,3]
func RightBoundSearch(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	left, right := 0, length-1
	for left <= right {
		mid := left + (right-left)/2
		value := nums[mid]
		switch {
		case value == target:
			// 扩大左侧边界
			left = mid + 1
		case value < target:
			left = mid + 1
		case value > target:
			right = mid - 1
		}
	}
	// TODO 注意边界条件
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}

// RightBoundSearchV1 优化边界条件
func RightBoundSearchV1(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	left, right := 0, length-1
	for left <= right {
		mid := left + (right-left)/2
		value := nums[mid]
		switch {
		case value == target:
			// 优化版边界条件
			if mid == length-1 || nums[mid+1] != target {
				return mid
			}
			// 扩大左侧边界
			left = mid + 1
		case value < target:
			left = mid + 1
		case value > target:
			right = mid - 1
		}
	}
	return right
}
