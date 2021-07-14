package bootcamp

/*
* [344] Reverse String
 */

// @lc code=start
func reverse(s []byte) {
	maxLen := len(s)
	for i := 0; i < maxLen/2; i++ {
		s[i], s[maxLen-i-1] = s[maxLen-i-1], s[i]
	}
}

func reverseV1(s []byte) {
	maxLen := len(s)
	// 双指针法
	left, right := 0, maxLen-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
