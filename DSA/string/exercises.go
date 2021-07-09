package string

/*
题目: Given a string, that contains a special character together with alphabets (‘a’ to ‘z’ and ‘A’ to ‘Z’), reverse the string in a way that special characters are not affected.
地址: https://www.geeksforgeeks.org/reverse-a-string-without-affecting-special-characters/
*/

func reverseStringWithoutAffectSpecialChar(str string) string {
	length := len(str)
	if length == 0 {
		return str
	}

	item := []byte(str)

	left, right := 0, length-1
	for left <= right {
		leftCharIsAlphabet := isAlphabet(item[left])
		rightCharIsAlphabet := isAlphabet(item[right])
		if !leftCharIsAlphabet {
			left++
		} else if !rightCharIsAlphabet {
			right--
		} else {
			//  左右两边都是字母时,且不是特殊字符时、
			swapByte(item, left, right)
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

func swapByte(items []byte, left, right int) {
	items[left], items[right] = items[right], items[left]
}
