package array

/*
奇偶排序，偶数在前，奇数在后
考察点：
	* 原地操作
*/
func evenOdd(input []int) []int {
	length := len(input)
	nextEven, nextOdd := 0, length-1
	for nextEven < nextOdd {
		if input[nextEven]%2 == 0 {
			nextEven++
		} else {
			input[nextEven], input[nextOdd] = input[nextOdd], input[nextEven]
			nextOdd--
		}
	}
	return input
}

/*
计算两个已排序数组的交集,
要求：
	1. 两个数组有重复元素时，交集中不允许出现重复数字

方法：
	* 暴力法 O(mn)
	* 如果两个数组长度相差较大，可以遍历较小的数组，此时时间复杂度为 O(mlogn), m 为较小的数组长度
	* 线性并行遍历 O(m+n)
	* 应用思想：
		- 不变量
*/

func intersect(a, b []int) []int {
	posA, posB := 0, 0
	intersection := []int{}

	for posA < len(a) && posB < len(b) {
		if a[posA] == b[posB] {
			// 避免重复元素多次追加
			if posA == 0 || a[posA] != a[posA-1] {
				intersection = append(intersection, a[posA])
			}
			posA++
			posB++
		} else if a[posA] > b[posB] {
			posB++
		} else if a[posA] < b[posB] {
			posA++
		}
	}

	return intersection
}

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

/*
删除数组中的元素，并返回新的数组长度
要求：
	- 原地删除
*/

func removeElement(items []int, target int) int {
	length := len(items)

	pos := 0
	for i := 0; i < length; i++ {
		if items[i] != target {
			items[pos] = items[i]
			pos++
		}
	}

	return pos
}

/*
上述代码存在以下问题：
	- 如遇到 [1，2，3，5，4] 删除 4 时，会产生不必要赋值的问题
	- 如遇到 [4, 1, 2, 3, 5] 删除 4 时似乎没必要左移 [1, 2, 3, 5]

当元素匹配到 target 时，我们可以将当前元素与最后一个元素进行交换，并释放最后一个元素。这实际上使数组的大小减少了 1。
请注意，被交换的最后一个元素可能是您想要移除的值。但是不要担心，在下一次迭代中，我们仍然会检查这个元素。
*/
func removeElementV1(items []int, target int) int {
	for i := 0; i < len(items); {
		if items[i] == target {
			items = append(items[:i], items[i+1:]...)
		} else {
			i++
		}
	}

	return len(items)
}

func removeElementV2(items []int, target int) int {
	length := len(items)
	left, right := 0, length-1

	pos := 0 // 记录重复元素移动的位置
	for left <= right {
		if items[left] == target {
			items[left], items[right] = items[right], items[left]
			right--
			pos++
		} else {
			left++
		}
	}

	return length - pos
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
