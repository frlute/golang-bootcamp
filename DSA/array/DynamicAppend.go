package array

/*
go 切片的动态扩容
使用 copy 因为它复制可以处理共享相同基础数组的源和目标片，正确地处理重叠片.
*/
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) {
		// 避免 n=0, allocate double what's needed, for future growth.
		newSlice := make([]byte, n, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}

	slice = slice[0:n]
	copy(slice[m:n], data)

	return slice
}
