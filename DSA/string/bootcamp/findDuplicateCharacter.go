package bootcamp

import "fmt"

/*
解题思路：哈希法，利用哈希特性，即 key 唯一
*/
func findDuplicateCharacter(item string) {
	if item == "" {
		return
	}
	cached := make(map[byte]int)
	for _, str := range []byte(item) {
		cached[str]++
	}

	for key, count := range cached {
		if count > 1 {
			// %q 表示输出带引号的字符
			fmt.Printf("%q:%d\n", key, count)
		}
	}
}
