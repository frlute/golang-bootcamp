package mysort

/*
插入算法的核心思想是取未排序区间中的元素，在已排序区间中找到合适的插入位置将其插入，
并保证已排序区间数据一直有序。重复这个过程，直到未排序区间中元素为空，算法结束。
步骤：
	1.从第一个元素开始，该元素可以认为已经被排序
	2.取出下一个元素，在已经排序的元素序列中从后向前扫描
	3.如果该元素（已排序）大于新元素，将该元素移到下一位置
	4.重复步骤3，直到找到已排序的元素小于或者等于新元素的位置
	5.将新元素插入到该位置后
	6.重复步骤2~5

按升序排序

时间复杂度 O(n), 空间复杂度 O(1) 也是一种分而治之的应用，数据分为已排序和未排序两部分，然后根据位置滑动比较
*/
func InsertionSort(items []int) {
	length := len(items)
	if length <= 1 {
		return
	}

	for index := 1; index < length; index++ {
		preIndex := index - 1
		value := items[index]
		// 寻找插入的位置
		for preIndex >= 0 && items[preIndex] > value {
			items[preIndex+1] = items[preIndex] // 数据移动
			preIndex--
		}

		//插入数据, 此时用 value 因如果数据移动，data 数据已变化
		if preIndex+1 != index {
			items[preIndex+1] = value
		}
	}
}

// InsertionSortV1 优化: 拆半插入
func InsertionSortV1(nums []int) {
	// length := len(nums)
	// if length <= 1 {
	// 	return
	// }
	// for i := 1; i < length; i++ {
	// 	prePos := i - 1
	// 	value := nums[i]
	// 	for prePos >= 0 && value < nums[prePos] {
	// 		helper.SwapIntArray(nums, prePos+1, prePos)
	// 		prePos--
	// 	}
	// }
}
