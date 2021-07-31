package tree

import (
	"testing"
)

func Test_BinarySearchTree_insert(t *testing.T) {
	rootTree := NewBinarySearchTree(8)
	rootTree.Insert(11)
	rootTree.Insert(2)
	// rootTree.Insert(1)
	// rootTree.Insert(7)
	// rootTree.Insert(8)
	rootTree.Display()
	// fmt.Println("----: ", rootTree.String())
}

func Test_BinarySearchTree_InOrderTraverse(t *testing.T) {
	rootTree := NewBinarySearchTree(8)
	rootTree.Insert(11)
	rootTree.Insert(2)
	rootTree.Display()
}

func Test_BinarySearchTree_PreOrderTraverse(t *testing.T) {
	rootTree := NewBinarySearchTree(8)
	rootTree.Insert(11)
	rootTree.Insert(2)
	rootTree.Display()
}

func Test_BinarySearchTree_PostOrderTraverse(t *testing.T) {
	rootTree := NewBinarySearchTree(8)
	rootTree.Insert(11)
	rootTree.Insert(2)
	rootTree.Display()
}

func Test_xx(t *testing.T) {
	// var a interface{} = 3
	// var b interface{} = 5
	// fmt.Println(b > a)
}
