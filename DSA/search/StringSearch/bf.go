package StringSearch

//bf BF(Brute Force) 算法, BF search pattern index, return the first match subs start index
func bf(src, target string) int {
	srcLength, targetLength := len(src), len(target)
	if srcLength == 0 || targetLength == 0 || srcLength < targetLength {
		return -1
	}

	for index := 0; index <= srcLength-targetLength; index++ {
		subStr := src[index : index+targetLength]
		if subStr == target {
			return index
		}
	}

	return -1
}
