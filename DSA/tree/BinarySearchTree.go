package tree

import "fmt"

// BinaryTreeBasedArray _
type binarySearchTree struct {
	root *BinaryTreeNode
}

// NewBinaryTreeBasedArray _
func NewBinarySearchTree(root interface{}) BinarySearchTree {
	return newBinarySearchTree(root)
}

func newBinarySearchTree(root interface{}) *binarySearchTree {
	rootNode := &BinaryTreeNode{Data: root}
	return &binarySearchTree{root: rootNode}
}

// TODO 存在 bug，insert 后根节点赋值存在问题
func (bst *binarySearchTree) Insert(element interface{}) {
	newNode := &BinaryTreeNode{Data: element}
	if bst.IsEmpty() {
		bst.root = newNode
	} else {
		currentNode := bst.root
		var parentNode *BinaryTreeNode

		for {
			parentNode = currentNode
			//go to left of the tree
			if element.(int) < currentNode.Data.(int) {
				currentNode = currentNode.Left
				//insert to the left node
				if currentNode == nil {
					parentNode.Left = newNode
					return
				}
			} else {
				currentNode = currentNode.Right
				// insert to the right node
				if currentNode == nil {
					parentNode.Right = newNode
					return
				}
			}
		}
	}
}

func (bst *binarySearchTree) Search(target interface{}) *BinaryTreeNode {
	if bst.IsEmpty() {
		return nil
	}

	currentNode := bst.root
	for currentNode.Data != target {
		if currentNode.Data != nil {
			fmt.Printf("node value: %+v", currentNode.Data)
		}
		if currentNode.Data.(int) < target.(int) {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}

		if currentNode.Data == nil {
			return nil
		}
	}

	return currentNode
}

func (bst *binarySearchTree) InOrderTraverse(node *BinaryTreeNode) {
	for !bst.IsEmpty() {
		fmt.Printf("%+v", node.Data)
		bst.InOrderTraverse(node.Left)
		bst.InOrderTraverse(node.Right)
	}
}

func (bst *binarySearchTree) PreOrderTraverse(node *BinaryTreeNode) {
	for !bst.IsEmpty() {
		bst.InOrderTraverse(node.Left)
		fmt.Printf("%+v", node.Data)
		bst.InOrderTraverse(node.Right)
	}
}

func (bst *binarySearchTree) PostOrderTraverse(node *BinaryTreeNode) {
	for !bst.IsEmpty() {
		bst.InOrderTraverse(node.Left)
		bst.InOrderTraverse(node.Right)
		fmt.Printf("%+v", node.Data)
	}
}

func (bst *binarySearchTree) Depth() int {
	return 0
}

func (bst *binarySearchTree) Level() int {
	return 0
}

func (bst *binarySearchTree) IsEmpty() bool {
	if bst.root == nil {
		return true
	}
	return false
}

func (bst *binarySearchTree) Display() {
	fmt.Printf("%+v\n", bst.root)
	// fmt.Println(bst.root.Left)
}

func (bst *binarySearchTree) String() string {
	// fmt.Printf("%+v", bst.root)
	return ""
}
