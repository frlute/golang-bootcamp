package stack

type Stack interface {
	// Pushing (storing) an element on the stack.
	Push(item interface{})
	Pop() interface{}

	// get the top data element of the stack, without removing it.
	Peek() interface{}
	IsFull() bool
	IsEmpty() bool

	Flush()
	Display()
}
