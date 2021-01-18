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
	quickSortHelper(data, 0, n-1)
}

func quickSortHelper(data []int, left, right int) {
	if left >= right {
		return
	}
	// 进行切分
	pivot := partition(data, left, right)
	// 对左部分进行快排
	quickSortHelper(data, left, pivot-1)
	// 对右部分进行快排
	quickSortHelper(data, pivot+1, right)
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
