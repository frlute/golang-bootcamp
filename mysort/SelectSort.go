package mysort

/*
选择排序算法的实现思路有点类似插入排序，也分已排序区间和未排序区间。
但是选择排序每次会从未排序区间中找到最小的元素，将其放到已排序区间的末尾。
*/
func SelectSort(data []int) {
	n := len(data)
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		// 假定首元素为最小元素
		min := i
		// 从i右侧的所有元素中找出当前最小值所在的下标
		for j := i + 1; j < n; j++ {
			if data[j] < data[min] {
				min = j
			}
		}
		// 把每次找出来的最小值与之前的最小值做交换
		if i != min {
			data[i], data[min] = data[min], data[i]
		}
	}
}
