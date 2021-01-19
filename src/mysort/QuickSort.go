package mysort

/*
快排的思想是这样的：
	如果要排序数组中下标从 p 到 r 之间的一组数据，
	我们选择 p 到 r 之间的任意一个数据作为 pivot（分区点）。
也是利用分治思想
*/
func QuickSort(data []int) {
	n := len(data)
	if n <= 1 {
		return
	}
	quickSort(data, 0, n-1)
}

func quickSort(data []int, left, right int) {
	if left >= right {
		return
	}
	// 进行切分
	pivot := partition(data, left, right)
	// 对左部分进行快排
	quickSort(data, left, pivot-1)
	// 对右部分进行快排
	quickSort(data, pivot+1, right)
}

func partition(list []int, left, right int) int {
	pivot := list[right]
	start := left - 1
	for i := left; i < right; i++ {
		if list[i] < pivot {
			start++
			list[i], list[start] = list[start], list[i]
		}
	}
	list[start+1], list[right] = list[right], list[start+1]

	return start + 1
}

// QuickSort2 xx
// O(nlogn) 稳定(待确认) 按小到大的方式排序, 最差运行时间O(n^2)
func QuickSort2(l []int) []int {
	length := len(l)
	// base case
	if length < 2 {
		return l
	}

	// select pivot
	// 初始下标为1，即基准值右侧的第一个元素的下标
	pivot, initIndex := l[0], 1

	// partitioning
	head, tail := 0, length-1
	for head < tail {
		// 如果基准值右侧的元素大于分水岭，就将右侧的尾部元素和分水岭右侧元素交换
		if l[initIndex] > pivot {
			l[initIndex], l[tail] = l[tail], l[initIndex]
			// 尾下标左移一位
			tail--
		} else {
			// 如果基准值右侧的元素小于等于分水岭，就将分水岭右移一位
			l[initIndex], l[head] = l[head], l[initIndex]
			// 头下标右移一位
			head++
			// 初始下标右移一位
			initIndex++
		}
	}

	// recursion
	QuickSort2(l[:head])
	QuickSort2(l[head+1:])

	return l
}
