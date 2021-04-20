package tree

// BinaryTree _
type BinaryTree struct {
	root *Node
}

// NewBinaryTree _
func NewBinaryTree(root interface{}) *BinaryTree {
	return &BinaryTree{NewNode(root)}
}

func (bt *BinaryTree) inOrderTraverse() {

}

func (bt *BinaryTree) preOrderTraverse() {

}

func (bt *BinaryTree) postOrderTraverse() {

}
