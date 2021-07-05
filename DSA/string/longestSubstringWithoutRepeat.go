package string

/*
题目：

要求：
	- 子字符串不能存在重复的字符

解法：
	- 收缩区间、双指针法
*/

func longestSubstringWithoutRepeat(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}

	cached := make(map[byte]int, length)

	left, right := 0, 0
	longest := 0
	for right < length {
		str := s[right]
		right++

		cached[str]++
		// 此时用 for 因不一定已跑过的字符是首个条件满足了，可能是最后一个满足了，故需要循环判断收缩 left 位置
		for cached[str] > 1 {
			d := s[left]
			left++

			cached[d]--
		}

		// 更新最长子字符
		if longest < (right - left) {
			longest = right - left
		}
	}

	return longest
}
