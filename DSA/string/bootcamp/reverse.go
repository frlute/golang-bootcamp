package bootcamp

/*
题目来源：https://javarevisited.blogspot.com/2012/01/how-to-reverse-string-in-java-using.html#axzz70sjwWbKO

解题思路：
	1.反向迭代法
	2.双指针法，类似迭代法
	3.递归法？
*/

func reverse(item string) string {
	length := len(item)

	list := []byte(item)
	left, right := 0, length-1
	for left <= right {
		list[left], list[right] = list[right], list[left]
		left++
		right--
	}
	return string(list)
}

// 递归法
func reverseV1(item string) string {
	if len(item) < 2 {
		return item
	}

	// TODO 需要意识到 string 其实是 []byte 类型的 slice
	return reverseV1(item[1:]) + item[:1]

}
