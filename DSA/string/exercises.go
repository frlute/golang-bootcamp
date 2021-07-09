package string

/*
题目: Given a string, that contains a special character together with alphabets (‘a’ to ‘z’ and ‘A’ to ‘Z’), reverse the string in a way that special characters are not affected.
地址: https://www.geeksforgeeks.org/reverse-a-string-without-affecting-special-characters/
*/

func reverseStingWithoutAffectSpecialChar(str string) string {
	length := len(str)
	if length == 0 {
		return str
	}

	item := []byte(str)

	left, right := 0, length-1
	for left <= right {
		leftCharIsAlphabet := isAlphabet(item[left])
		rightCharIsAlphabet := isAlphabet(item[right])
		if leftCharIsAlphabet && rightCharIsAlphabet {
			// 左右两边都是字母时
			item[left], item[right] = item[right], item[left]
			left++
			right--
		} else if leftCharIsAlphabet && !rightCharIsAlphabet {
			// 左边是字母右边不是时
			right--
		} else if !leftCharIsAlphabet && rightCharIsAlphabet {
			// 左边不是字母右边是字母
			left--
		} else {
			// 两边都不是字母
			left++
			right--
		}
	}
	return string(item)
}

func isAlphabet(item byte) bool {
	isUpperAlphabet := item >= 'A' && item <= 'Z'
	isLowerAlphabet := item >= 'a' && item <= 'z'
	return isUpperAlphabet || isLowerAlphabet
}
