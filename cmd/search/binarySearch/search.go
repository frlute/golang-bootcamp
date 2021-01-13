package binarySearch

func Search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	left, right := 0, length-1
	// 必须按顺序排序时才能这样判断
	if nums[right] > target {
		return -1
	}

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

// LeftBoundSearch 表示从左边界开始搜索, such as [1,2,2,2,3]
func LeftBoundSearch(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	left, right := 0, length
	// 必须按顺序排序时才能这样判断
	// if nums[right-1] > target {
	// 	return -1
	// }
	for left < right {
		mid := (left + right) / 2
		value := nums[mid]
		if value == target {
			// 此时并不返回
			right = mid
		} else if value < target {
			left = mid + 1 // 注意
		} else if value > target {
			right = mid // 注意
		}
	}

	if left == length {
		return -1
	}
	if nums[left] == target {
		return left
	}

	return -1
}
