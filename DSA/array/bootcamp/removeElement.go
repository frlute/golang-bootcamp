package bootcamp

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
