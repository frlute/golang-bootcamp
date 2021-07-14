package bootcamp

import (
	"math"
	"strings"
)

/*
string 转换为数字, 支持负数
提示： 一次生成一个数字的结果
solution:
	- 多个字符时逐位转换
	- 注意正负数符号
	- 注意转换后的百分位，
*/

func convertToInt(item string) int {
	length := len(item)
	if length <= 0 {
		panic("this string is empty value.")
	}
	isNegative := false
	if strings.HasPrefix(item, "-") {
		isNegative = true
		item = item[1:]
		length--
	}

	var count int
	for i := 0; i < length; i++ {
		// byte, utf-8 转成真实的数字？
		d := item[i] - '0'
		//  10^2 × d2 + 10^1 × d1 + d0
		count += int(d) * int(math.Pow10(length-i-1))
	}

	if isNegative {
		return -count
	}
	return count
}

// 以 341 为例，3， 3*10+4， 34*10+1
func convertToIntV2(item string) int {
	length := len(item)
	if length <= 0 {
		panic("this string is empty value.")
	}
	isNegative := false
	if strings.HasPrefix(item, "-") {
		isNegative = true
		item = item[1:]
		length--
	}

	var count int
	for i := 0; i < length; i++ {
		// byte, utf-8 转成真实的数字？
		d := item[i] - '0'
		if count == 0 {
			count = int(d)
		} else {
			count = count*10 + int(d)
		}
	}

	if isNegative {
		return -count
	}
	return count
}
