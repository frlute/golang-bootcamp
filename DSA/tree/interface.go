package tree

type BinarySearchTree interface {
	Insert(element interface{})
	// return the target node
	Search(target interface{}) *BinaryTreeNode
	PreOrderTraverse(node *BinaryTreeNode)
	InOrderTraverse(node *BinaryTreeNode)
	PostOrderTraverse(node *BinaryTreeNode)

	Depth() int
	Level() int

	IsEmpty() bool
	Display()
	String() string
}

type ALVTree interface {
}
