package tree

type BinaryTree interface {
	PreOrderTraverse()
	InOrderTraverse()
	PostOrderTraverse()

	Depth() int64
}
