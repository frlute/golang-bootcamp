package mysort

/*
选择排序算法的实现思路有点类似插入排序，也分已排序区间和未排序区间。
但是选择排序每次会从未排序区间中找到最小的元素，将其放到已排序区间的末尾。
*/
func SelectSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}

	for index := 0; index < length; index++ {
		// 假定首元素为最小元素
		min := index
		// 从 index 右侧的所有元素中找出当前最小值所在的下标
		for j := index + 1; j < length; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		// 把每次找出来的最小值与之前的最小值做交换
		if index != min {
			nums[index], nums[min] = nums[min], nums[index]
		}
	}
}
