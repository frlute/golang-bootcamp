package binarySearch

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
			// 此时并不返回, 不断向 left 递进查到最左侧的值
			right = mid
		} else if value < target {
			left = mid + 1 // 注意
		} else if value > target {
			right = mid // 注意
		}
	}

	// target 比所有数都大
	if left == length {
		return -1
	}

	result := func(left int) int {
		if nums[left] == target {
			return left
		}
		return -1
	}

	return result(left)
}

// LeftBoundSearchV1 表示优化版本
func LeftBoundSearchV1(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	left, right := 0, length-1
	for left <= right {
		mid := (left + (right - left)) / 2
		value := nums[left]

		switch {
		case value == target:
			// 收缩右侧边界
			right = mid - 1
		case value < target:
			// 搜索区间变为 [mid+1, right]
			left = mid + 1
		case value > target:
			// 搜索区间变为 [left, mid-1]
			right = mid - 1
		}
	}

	// 检查出界情况
	if left > length || nums[left] != target {
		return -1
	}

	return left
}
