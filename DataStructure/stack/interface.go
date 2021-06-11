package stack

type Stack interface {
	Push(item interface{}) bool
	Pop() interface{}
	Len() int64
	Cap() int64
}
