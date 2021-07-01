package array

/*
奇偶排序，偶数在前，奇数在后
考察点：
	* 原地操作
*/
func evenOdd(input []int) []int {
	length := len(input)
	nextEven, nextOdd := 0, length-1
	for nextEven < nextOdd {
		if input[nextEven]%2 == 0 {
			nextEven++
		} else {
			input[nextEven], input[nextOdd] = input[nextOdd], input[nextEven]
			nextOdd--
		}
	}
	return input
}
