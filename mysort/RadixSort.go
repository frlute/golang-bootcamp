package mysort

/*
基数排序对要排序的数据是有要求的，需要可以分割出独立的“位”来比较，而且位之间有递进的关系，
如果 a 数据的高位比 b 数据大，那剩下的低位就不用比较了。除此之外，每一位的数据范围不能太大，
要可以用线性排序算法来排序，否则，基数排序的时间复杂度就无法做到 O(n) 了。

步骤：
1.将所有待比较数值（正整数）统一为同样的数位长度，数位较短的数前面补零
2.从最低位开始，依次进行一次排序
3.从最低位排序一直到最高位排序完成以后, 数列就变成一个有序序列
*/
func RadixSort(data []int) {
	length := len(data)
	if length <= 1 {
		return
	}

	// 获取最大位数
	maxSize := maxBit(data)

	tmp := make([]int, length)

	buckets := new([10]int)
}

func maxBit(data []int) int {
	ret := 1
	maxValue := 10
	for _, value := range data {
		if value > maxValue {
			maxValue = value * 10
			ret++
		}
	}

	return ret
}
