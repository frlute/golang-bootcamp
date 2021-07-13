package mysort

/*
归并排序的核心思想还是蛮简单的。如果要排序一个数组，我们先把数组从中间分成前后
两部分，然后对前后两部分分别排序，再将排好序的两部分合并在一起，这样整个数组就都有序了。
使用分治思想
分治算法一般都是用递归来实现的。分治是一种解决问题的处理思想，递归是一种编程技巧
*/
func MergeSort(data []int) []int {
	length := len(data)
	if length <= 1 {
		return data
	}

	mid := length / 2
	// 分治递归
	left := MergeSort(data[0:mid])
	right := MergeSort(data[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}

	// [左右]对比，是指左的第一个元素，与右边的第一个元素进行对比，哪个小，
	// 就先放到结果的第一位，然后左或右取出了元素的那边的索引进行
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	// 比较完后，还要分别将左，右的剩余的元素，追加到结果列的后面(不然就漏咯）
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}

// TODO 可以用哨兵方法优化 merge 函数
