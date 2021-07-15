package bootcamp

/*
问题来源： https://www.hackerrank.com/challenges/array-left-rotation/problem

* Complete the 'rotateLeft' function below.
*
* The function is expected to return an INTEGER_ARRAY.
* The function accepts following parameters:
*  1. INTEGER d
*  2. INTEGER_ARRAY arr
*/

func rotateLeft(d int32, arr []int32) []int32 {
	// Write your code here
	length := len(arr)
	if length <= 1 {
		return arr
	}
	res := make([]int32, 0, len(arr))
	res = append(res, arr[d:]...)
	res = append(res, arr[:d]...)

	return res
}

// 题解中提供的迭代解法, 此种方法更容易理解一些
func rotateLeftV1(d int32, arr []int32) []int32 {
	length := int32(len(arr))
	if length <= 1 {
		return arr
	}

	res := make([]int32, len(arr))
	var index int32
	for ; index < length; index++ {
		pos := (index - d) % length
		if pos < 0 {
			pos += length
		}
		res[pos] = arr[index]
	}
	return res
}
