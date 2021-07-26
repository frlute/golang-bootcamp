package array

// Array 自己实现 array 常用的操作, 用 []int 来示例
type Array interface {
	Append(value ...int) []int
	Insert(index, value int)

	Cap() int
	Len() int

	GetIndex(value int) int
	Remove(value int)

	Display()
}
