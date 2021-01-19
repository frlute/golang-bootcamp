package mysort

/*
桶排序，顾名思义，会用到“桶”，核心思想是将要排序的数据分到几个有序的桶里，
每个桶里的数据再单独进行排序。桶内排完序之后，再把每个桶里的数据按照顺序依次取出，
组成的序列就是有序的了。

桶排序比较适合用在外部排序中。所谓的外部排序就是数据存储在外部磁盘中，数据量比较大，
内存有限，无法将数据全部加载到内存中。

桶排序以下列程序进行：

1.设置一个定量的数组当作空桶子。
2.寻访序列，并且把项目一个一个放到对应的桶子去。
3.对每个不是空的桶子进行排序。
4.从不是空的桶子里把项目再放回原来的序列中
*/
func BucketSort(data []int) {
	length := len(data)
	//k（数组最大值）
	max := getMaxInArr(data)
	//二维切片
	buckets := make([][]int, length)

	//分配入桶
	index := 0
	for i := 0; i < length; i++ {
		//分配桶 index = value * (n-1) /k
		index = data[i] * (length - 1) / max
		buckets[index] = append(buckets[index], data[i])
	}

	// 桶内排序
	tmpPos := 0
	for j := 0; j < length; j++ {
		bucketLen := len(buckets[j])
		if bucketLen > 0 {
			InsertionSort(buckets[j])
			copy(data[tmpPos:], buckets[j])
			tmpPos += bucketLen
		}
	}
}

func getMaxInArr(data []int) int {
	max := 0
	for _, value := range data {
		if value > 0 {
			max = value
		}
	}
	return max
}
