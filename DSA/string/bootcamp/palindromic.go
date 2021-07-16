package bootcamp

import "fmt"

/*
判断字符串是否回文字符串
- 方法一: 创建一个新字符串，反向遍历，然后与原字符串对比，此时空间复杂为 O(n), 时间复杂度为 O(n)
- 方法二：遍历字符串，对比第 i 个字符与第 n-i 个字符相同，相同则为回文字符串，此方式空间复杂度为 O(0), 时间复杂度为 O(n)
*/

func isPalindromic(item string) bool {
	length := len(item)
	if length <= 1 {
		return false
	}

	// TODO 当 length 为偶数时，会多遍历一次
	loop := 0
	for i := 0; i < length/2; i++ {
		loop++
		if item[i] != item[length-i-1] {
			return false
		}
	}

	fmt.Println("v0 loop: ", loop)
	return true
}

// 前后指针法
func isPalindromicV1(item string) bool {
	length := len(item)
	if length <= 1 {
		return false
	}

	left, right := 0, length-1
	loop := 0
	for left < right {
		loop++
		if item[left] != item[right] {
			return false
		}
		left++
		right--

	}
	fmt.Println("v1 loop: ", loop)

	return true
}

// TODO 回文相关题目还有很多变种，待练习
