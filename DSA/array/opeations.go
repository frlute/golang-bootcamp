package array

type array = []int

// func newArray() Array {
// 	return &array{}
// }

// ContainsInt 最好通过 map 方式判断是否存在
func ContainsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ContainsRune _
func ContainsByte(s []byte, e byte) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Delete(items []byte, target byte) []byte {
	pos := Index(items, target)
	if pos != -1 {
		result := make([]byte, 0, len(items)-1)
		result = append(result, items[:pos]...)
		result = append(result, items[pos+1:]...)

		return result
	}
	return items
}

func Index(items []byte, target byte) int {
	for index, value := range items {
		if value == target {
			return index
		}
	}
	return -1
}

func RemoveAt(items []rune, pos int) []rune {
	result := make([]rune, 0, len(items)-1)
	result = append(result, items[:pos]...)
	result = append(result, items[pos+1:]...)

	return result
}
