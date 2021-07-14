package bootcamp

import "fmt"

/*
* [4] Median of Two Sorted Arrays

思路：
	1. 使用归并的方式，合并两个有序数组，得到一个大的有序数组。大的有序数组的中间位置的元素，即为中位数。
	2. 不需要合并两个有序数组，只要找到中位数的位置即可。由于两个数组的长度已知，因此中位数对应的两个数组的下标之和也是已知的。维护两个指针，初始时分别指向两个数组的下标 00 的位置，每次将指向较小值的指针后移一位（如果一个指针已经到达数组末尾，则只需要移动另一个数组的指针），直到到达中位数的位置。
*/
// 时间复杂度 O(m+n), 且空间复杂度为 O(m+n), 不符合题目时间复杂度为 O(log(m+n)) 要求
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)

	pos1, pos2 := 0, 0

	result := make([]int, 0, l1+l2)
	loop := 1
	for pos1 < l1 && pos2 < l2 {
		if nums1[pos1] > nums2[pos2] {
			result = append(result, nums2[pos2])
			pos2++
		} else if nums1[pos1] < nums2[pos2] {
			result = append(result, nums1[pos1])
			pos1++
		} else {
			result = append(result, nums1[pos1], nums2[pos2])
			pos1++
			pos2++
		}
		loop++
	}

	// num1
	if pos1 <= l1 {
		result = append(result, nums1[pos1:]...)
	}
	if pos2 <= l2 {
		result = append(result, nums2[pos2:]...)
	}

	fmt.Println(result)

	length := len(result)
	midPos := length / 2
	if length%2 == 0 {
		return float64(result[midPos-1]+result[midPos]) / 2
	}

	return float64(result[midPos])
}

// TODO 待实现二分查找的优化实现
func findMedianSortedArraysV1(nums1 []int, nums2 []int) float64 {
	return float64(0)
}
