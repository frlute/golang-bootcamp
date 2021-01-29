package binarySearch

// LastLeBSearch last less than or equal
func LastLeBSearch(items []int, target int) int {
	length := len(items)
	if length == 0 {
		return -1
	}
	low, high := 0, length-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		value := items[mid]
		if value <= target {
			if mid == length-1 || items[mid+1] > target {
				return mid
			}
			low = mid + 1
		} else if value > target {
			high = mid - 1
		}
	}
	return -1
}
