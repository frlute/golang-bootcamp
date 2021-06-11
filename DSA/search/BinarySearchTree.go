package search

import (
	"github.com/frlute/golang-bootcamp/DSA/tree"
)

//BST BinarySearchTree
type BST struct {
	// data  int
	// left  *BST
	// right *BST
	*tree.BinaryTreeBasedArray
}

// NewBST _
func NewBST(item int) *BST {
	return &BST{
		BinaryTreeBasedArray: tree.NewBinaryTreeBasedArray(item),
	}
}

func (b *BST) find(target int) *tree.Node {
	tmp := b.Root
	for tmp != nil {
		value := tmp.Data.(int)
		switch {
		case value > target:
			tmp = tmp.Left
		case value < target:
			tmp = tmp.Right
		case value == target:
			return tmp
		}
	}
	return nil
}

// 新插入的数据一般都是在叶子节点上
func (b *BST) insert(item int) bool {
	if b.Root == nil {
		b.Root = tree.NewNode(item)
		return true
	}

	tmp := b.Root
	for tmp != nil {
		value := tmp.Data.(int)
		if item == value {
			return false
		} else if item < value {
			if tmp.Left == nil {
				tmp.Left = tree.NewNode(item)
				break
			}
			tmp = tmp.Left
		} else {
			if tmp.Right == nil {
				tmp.Right = tree.NewNode(item)
				break
			}
			tmp = tmp.Right
		}
	}

	return true
}

func (b *BST) delete(target int) {
	if b.Root == nil {
		return
	}

	// 查询删除的节点
	deleteNode := b.Root
	var parentNode *tree.Node // 要删除节点的父节点
	for deleteNode != nil && deleteNode.Data != target {
		value := deleteNode.Data.(int)
		parentNode = deleteNode
		if target > value {
			deleteNode = deleteNode.Right
		} else if target < value {
			deleteNode = deleteNode.Left
		}
	}
	// 没找到删除的节点
	if deleteNode == nil {
		return
	}

	// 删除的节点同时存在左右节点时
	if deleteNode.Left != nil && deleteNode.Right != nil {
		minNode := deleteNode.Right
		minParentNode := deleteNode
		for minNode.Left != nil {
			minParentNode = minNode
			minNode = minNode.Left
		}

		// 将 minNode 的数据替换到 b 中
		deleteNode.Data = minNode.Data
		// 下面就变成了删除 minNode 了
		deleteNode = minNode
		parentNode = minParentNode
	}

	// 删除节点是叶子节点或者仅有一个子节点
	var child *tree.Node // deleteNode 的子节点
	if deleteNode.Left != nil {
		child = deleteNode.Left
	} else if deleteNode.Right != nil {
		child = deleteNode.Right
	} else {
		child = nil
	}

	if parentNode == nil {
		// 删除的是根节点时
		b.Root = child
	} else if parentNode.Left == deleteNode {
		parentNode.Left = child
	} else {
		parentNode.Right = child
	}
}

// 获取值最小节点
func (b *BST) min() *tree.Node {
	node := b.Root
	for node.Left != nil {
		node = node.Left
	}

	return node
}

// 获取值最大节点
func (b *BST) max() *tree.Node {
	node := b.Root
	for node.Right != nil {
		node = node.Right
	}

	return node
}
