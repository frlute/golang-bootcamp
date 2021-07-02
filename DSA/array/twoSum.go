package array

/*
两数之和
要求：
	- 已排序的整型数组
	- 只要满足两数之和即可
*/

// 暴力法 时间复杂度: O(n^2)
func twoSum(items []int, target int) []int {
	length := len(items)
	if length <= 1 {
		return nil
	}

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if items[i]+items[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

// 哈希法 时间复杂度 O(n), 空间复杂度 O(n), 应用以时间换空间的思想
func twoSumBasedHash(items []int, target int) []int {
	length := len(items)
	if length <= 1 {
		return nil
	}
	cached := make(map[int]int, length)
	for index, value := range items {
		expectValue := target - value
		if idx, ok := cached[expectValue]; ok {
			return []int{idx, index}
		}
		cached[value] = index
	}

	return nil
}

// 利用不变量思想、收缩范围，此种解法需要数组已排序, 时间复杂度 O(n)
func twoSumBasedInvariants(items []int, target int) []int {
	length := len(items)
	if length <= 1 {
		return nil
	}

	left, right := 0, length-1
	for left <= right {
		value := items[left] + items[right]
		if value == target {
			return []int{left, right}
		} else if value > target {
			right--
		} else {
			left++
		}
	}

	return nil
}
