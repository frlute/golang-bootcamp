package tree

import "fmt"

// Node _
type Node struct {
	Data  interface{}
	Left  *Node
	Right *Node
}

// NewNode _
func NewNode(data interface{}) *Node {
	return &Node{
		Data: data}
}

// String _
func (n *Node) String() string {
	return fmt.Sprintf("v:%+v, left:%+v, right:%+v", n.Data, n.Left, n.Right)
}
