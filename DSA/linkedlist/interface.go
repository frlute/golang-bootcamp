package linkedlist

// LinkedList define the methods provided by the linked list
type LinkedList interface {
	// adds an item to the end of the linked list
	Append(value interface{})
	Prepend(value interface{})
	Insert(value interface{}, pos int) error

	RemoveAt(pos int)
	Delete(value interface{})

	IndexOf(value interface{}) int

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
