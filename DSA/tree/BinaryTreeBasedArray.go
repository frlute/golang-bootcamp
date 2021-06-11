package tree

// BinaryTreeBasedArray _
type BinaryTreeBasedArray struct {
	Root *Node
}

// NewBinaryTreeBasedArray _
// func NewBinaryTreeBasedArray(root interface{}) BinaryTree {
// 	return newBinaryTreeBasedArray(root)
// }

func NewBinaryTreeBasedArray(root interface{}) *BinaryTreeBasedArray {
	return &BinaryTreeBasedArray{NewNode(root)}
}

func (bt *BinaryTreeBasedArray) InOrderTraverse() {

}

func (bt *BinaryTreeBasedArray) PreOrderTraverse() {

}

func (bt *BinaryTreeBasedArray) PostOrderTraverse() {

}

func (bt *BinaryTreeBasedArray) Depth() int64 {
	return 0
}
