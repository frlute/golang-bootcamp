package tree

import "fmt"

// Node _
type Node struct {
	data  interface{}
	left  *Node
	right *Node
}

// NewNode _
func NewNode(data interface{}) *Node {
	return &Node{data: data}
}

// String _
func (n *Node) String() string {
	return fmt.Sprintf("v:%+v, left:%+v, right:%+v", n.data, n.left, n.right)
}
