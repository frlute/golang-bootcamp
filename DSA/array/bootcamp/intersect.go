package bootcamp

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
