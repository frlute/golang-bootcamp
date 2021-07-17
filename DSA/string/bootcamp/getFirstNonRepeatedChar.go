package bootcamp

import (
	"strings"

	"github.com/frlute/golang-bootcamp/DSA/array"
)

/*
题目来源：https://javarevisited.blogspot.com/2014/03/3-ways-to-find-first-non-repeated-character-String-programming-problem.html

解题思路：
	1. 哈希法记录每个字符出现的数量，然后遍历
	2. 双层迭代法，i := 0 和 j := i+1

*/

// 哈希法，时间复杂度 O(2n), 空间复杂度 O(n)
func getFirstNonRepeatedChar(item string) string {
	length := len(item)
	if length <= 1 {
		return item
	}

	cached := make(map[byte]int)
	for _, str := range []byte(item) {
		cached[str]++
	}

	for _, value := range []byte(item) {
		if count, ok := cached[value]; ok && count == 1 {
			return string(value)
		}
	}

	return ""
}

func getFirstNonRepeatedCharV1(item string) string {
	length := len(item)
	if length <= 1 {
		return item
	}
	var repeating string
	nonRepeating := make([]byte, 0)

	for _, str := range []byte(item) {
		if strings.ContainsRune(repeating, rune(str)) {
			continue
		} else if array.ContainsByte(nonRepeating, str) {
			nonRepeating = array.Delete(nonRepeating, str)
			repeating += string(str)
		} else {
			nonRepeating = append(nonRepeating, str)
		}
	}
	if len(nonRepeating) > 0 {
		return string(nonRepeating[0])
	}
	return ""
}
