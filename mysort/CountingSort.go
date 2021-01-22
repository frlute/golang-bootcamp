package mysort

/*
计数排序可以理解为桶排序的一种特殊情况。
和桶排序不同的是计数排序中桶存的不是元素而是元素的数量。

计数排序只能用在数据范围不大的场景中，如果数据范围 k 比要排序的数据 n 大很多，
就不适合用计数排序了。而且，计数排序只能给非负整数排序，如果要排序的数据是其他类型的，
要将其在不改变相对大小的情况下，转化为非负整数

步骤：
1.找出待排序的数组中最大和最小的元素
2.统计数组中每个值为 i的元素出现的次数，存入数组 C 的第i项
3.对所有的计数累加 C 中的第一个元素开始，每一项和前一项相加）
4.反向填充目标数组：将每个元素i放在新数组的第C[i]项，每放一个元素就将C[i]减去1
*/
func CountingSort(data []int) []int {
	length := len(data)
	if length <= 1 {
		return data
	}
	max, min := getMaxAndMInInArr(data)
	bucketLen := max - min + 1
	buckets := make([]int, bucketLen)

	// 计算每个元素的个数，放入c中
	for i := range data {
		buckets[data[i]]++
	}

	result := make([]int, length)
	index := 0
	for k, v := range buckets {
		kk := k + min
		for j := v; j > 0; j-- {
			result[index] = kk
			index++
		}
	}

	return result
}

// CountingSort2 改进版
func CountingSort2(data []int) {
	length := len(data)
	if length <= 1 {
		return
	}
	max, _ := getMaxAndMInInArr(data)
	buckets := make([]int, max+1)

	// 计算每个元素的个数，放入c中
	for i := range data {
		buckets[data[i]]++
	}

	// 依次累加
	for i := 1; i <= max; i++ {
		buckets[i] = buckets[i-1] + buckets[i]
	}

	// 临时数组r，存储排序之后的结果
	result := make([]int, length)
	// 计算排序的关键步骤，有点难理解
	for i := length - 1; i >= 0; i-- {
		index := buckets[data[i]] - 1
		result[index] = data[i]
		buckets[data[i]]--
	}

	// 将结果拷贝给a数组
	for i := 0; i < length; i++ {
		data[i] = result[i]
	}
}

func getMaxAndMInInArr(data []int) (int, int) {
	max, min := data[0], data[0]
	for _, value := range data {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}
	return max, min
}
