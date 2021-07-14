package bootcamp

import (
	"fmt"
	"strings"
)

/*
判断两个字符串是否异构字符

这个题目理解起来有点绕，刚开始看了好一会都没理解意思~

解题思路: 只要俩字符串长度相同，且每个字符所对应的相对位置相同即可

考察点：
	- 对字符串的理解，同时也考察了 go 中字符串的相关操作及对字符串的理解
*/

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		fmt.Println(strings.IndexByte(s, s[i]), strings.IndexByte(t, t[i]))
		if strings.IndexByte(s, s[i]) != strings.IndexByte(t, t[i]) {
			return false
		}
	}

	return true
}

func isIsomorphicBasedHash(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[byte]int, len(s))
	tMap := make(map[byte]int, len(t))

	for i := range s {
		a1, ok1 := sMap[s[i]]
		a2, ok2 := tMap[t[i]]

		if ok1 && ok2 {
			if a1 != a2 {
				return false
			}
		} else if !ok1 && !ok2 {
			sMap[s[i]] += i
			tMap[t[i]] += i
		} else {
			return false
		}

	}

	return true
}
