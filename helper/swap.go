package helper

func SwapIntArray(items []int, left, right int) {
	items[left], items[right] = items[right], items[left]
}
