package linkedlist

// LinkedList define the methods provided by the linked list
type LinkedList interface {
	// adds an item to the end of the linked list
	Append(value interface{})
	Prepend(value interface{})
	Insert(value interface{}, pos int) error

	RemoveAt(pos int)
	Delete(value interface{})

	// The indexOf() method returns the first index at which a given element can be found in the array, or -1 if it is not present.
	IndexOf(value interface{}) int
	// 根据值搜索对应的节点
	Search(value interface{}) interface{}

	IsEmpty() bool
	// get the linked list size
	Size() int

	// display a string representation of the list
	Display()

	// returns a string representation of the list
	String() string

	// reset linked list
	Reset()
}
