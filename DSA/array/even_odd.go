package array

/*
奇偶排序，奇数在前，偶尔在后

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
