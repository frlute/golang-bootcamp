package mysort

/*
ShellSort _
希尔排序基于插入排序
希尔排序是基于插入排序的以下两点性质而提出改进方法的：
- 插入排序在对几乎已经排好序的数据操作时，效率高，即可以达到线性排序的效率；
- 但插入排序一般来说是低效的，因为插入排序每次只能将数据移动一位；

希尔排序的基本思想是：先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序，待整个序列中的记录"基本有序"时，再对全体记录进行依次直接插入排序。

相比于插入排序避免了大位移，主要是通过较小的值在最右侧时必须移到最左侧
利用 Knuth's Formula h=h*3+1, 把数据拆分多个单元间隔
最后利用插入排序排序整个数组

时间复杂度 O(n), 非稳定排序
*/
func ShellSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	interval := 1 // 或 gap
	if interval < length/3 {
		// Knuth's Formula
		interval = interval*3 + 1
	}
	// 记录迭代次数
	count := 0
	for interval > 0 {
		// fmt.Println("iteration count: ", count)
		for i := interval; i < length; i++ {
			valueToInsert := nums[i]
			pos := i

			for pos > interval-1 && nums[pos-interval] >= valueToInsert {
				nums[pos] = nums[pos-interval]
				pos -= interval
			}

			if i != pos {
				nums[pos] = valueToInsert
				// fmt.Printf("item inserted :%d, at position :%d\n", valueToInsert, pos)
			}
		}
		interval = (interval - 1) / 3
		count++
	}
}
