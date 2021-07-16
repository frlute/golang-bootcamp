package bootcamp

import (
	"sort"
	"strings"
)

/*
	- 易位构词 问题，看解释 https://zh.wikipedia.org/wiki/%E6%98%93%E4%BD%8D%E6%9E%84%E8%AF%8D%E6%B8%B8%E6%88%8F
	- String Anagram Example.
	- checks if two Strings are anagrams or not

想到的解法：
	1. 哈希，比对长度及每个字符数量是否一致
	2. 排序后字符串相等
	3.

*/
func isAnagram(a, b string) bool {
	lengthA := len(a)
	lengthB := len(b)
	if lengthA != lengthB {
		return false
	}

	if sortString(a) == sortString(b) {
		return true
	}

	return false
}

// 假定a, b 字符相同且都是小写
func isAnagramV1(a, b string) bool {
	lengthA := len(a)
	lengthB := len(b)
	if lengthA != lengthB {
		return false
	}

	for _, str := range []byte(a) {
		index := strings.IndexByte(b, str)
		if index != -1 {
			// string substring
			tmp := make([]byte, 0)
			tmp = append(tmp, []byte(b)[:index]...)
			tmp = append(tmp, []byte(b)[index+1:]...)
			b = string(tmp)
		} else {
			return false
		}

	}
	return b == ""
}

func isAnagramV2(a, b string) bool {
	lengthA := len(a)
	lengthB := len(b)
	if lengthA != lengthB {
		return false
	}

	mapA := make(map[byte]int)
	mapB := make(map[byte]int)
	for _, str := range []byte(a) {
		mapA[str]++
	}
	for _, str := range []byte(a) {
		mapB[str]++
	}
	for key, count := range mapA {
		num, ok := mapB[key]
		if !ok || num != count {
			return false
		}
	}

	return true
}

func sortString(item string) string {
	s := strings.Split(item, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func sortStringV1(item string) string {
	s := []rune(item)
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return string(s)
}

// TODO 易位构词还有很多变种题目，可以多了解下
