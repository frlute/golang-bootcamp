package tree

import (
	_ "github.com/frlute/go-playground/DSA/linkedlist"
)

// NodeBaseOnLinked _
type NodeBaseOnLinked struct {
	data  interface{}
	left  *NodeBaseOnLinked
	right *NodeBaseOnLinked
}

func (t *NodeBaseOnLinked) preOrder() {
	// if root == null {
	// 	return
	// }
	// print root // 此处为伪代码，表示打印root节点
	// preOrder(root->left)
	// preOrder(root->right)
}

func (t *NodeBaseOnLinked) inOrder() {
	// if root == null {
	// 	return
	// }
	// preOrder(root->left)
	// print root // 此处为伪代码，表示打印root节点
	// preOrder(root->right)
}

func (t *NodeBaseOnLinked) postOrder() {
	// if root == null {
	// 	return
	// }
	// preOrder(root->left)
	// preOrder(root->right)
	// print root // 此处为伪代码，表示打印root节点
}
