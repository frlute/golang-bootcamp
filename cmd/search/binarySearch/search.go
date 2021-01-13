package binarySearch

// Search 二分查找
func Search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	left, right := 0, length-1
	// 必须按顺序排序时才能这样判断
	// if nums[right] < target {
	// 	return -1
	// }

	for left <= right {
		// mid := (left + right) / 2
		// 需要防止溢出，代码中 left + (right - left) / 2 就和 (left + right) / 2 的结果相同，但是有效防止了 left 和 right 太大直接相加导致溢出。
		mid := left + (right-left)/2
		value := nums[mid]
		switch {
		case value == target:
			return mid
		case value < target:
			left = mid + 1
		case value > target:
			right = mid - 1
		}
	}

	return -1
}
