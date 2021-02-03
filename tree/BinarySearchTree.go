package tree

//BST BinarySearchTree
type BST struct {
	// data  int
	// left  *BST
	// right *BST
	*BinaryTree
}

// NewBST _
func NewBST(item int) *BST {
	return &BST{
		BinaryTree: NewBinaryTree(item),
	}
}

func (b *BST) find(target int) *Node {
	tmp := b.root
	for tmp != nil {
		value := tmp.data.(int)
		switch {
		case value > target:
			tmp = tmp.left
		case value < target:
			tmp = tmp.right
		case value == target:
			return tmp
		}
	}
	return nil
}

// 新插入的数据一般都是在叶子节点上
func (b *BST) insert(item int) bool {
	if b.root == nil {
		b.root = NewNode(item)
		return true
	}

	tmp := b.root
	for tmp != nil {
		value := tmp.data.(int)
		if item == value {
			return false
		} else if item < value {
			if tmp.left == nil {
				tmp.left = NewNode(item)
				break
			}
			tmp = tmp.left
		} else {
			if tmp.right == nil {
				tmp.right = NewNode(item)
				break
			}
			tmp = tmp.right
		}
	}

	return true
}

func (b *BST) delete(target int) {
	if b.root == nil {
		return
	}

	// 查询删除的节点
	deleteNode := b.root
	var parentNode *Node // 要删除节点的父节点
	for deleteNode != nil && deleteNode.data != target {
		value := deleteNode.data.(int)
		parentNode = deleteNode
		if target > value {
			deleteNode = deleteNode.right
		} else if target < value {
			deleteNode = deleteNode.left
		}
	}
	// 没找到删除的节点
	if deleteNode == nil {
		return
	}

	// 删除的节点同时存在左右节点时
	if deleteNode.left != nil && deleteNode.right != nil {
		minNode := deleteNode.right
		minParentNode := deleteNode
		for minNode.left != nil {
			minParentNode = minNode
			minNode = minNode.left
		}

		// 将 minNode 的数据替换到 b 中
		deleteNode.data = minNode.data
		// 下面就变成了删除 minNode 了
		deleteNode = minNode
		parentNode = minParentNode
	}

	// 删除节点是叶子节点或者仅有一个子节点
	var child *Node // deleteNode 的子节点
	if deleteNode.left != nil {
		child = deleteNode.left
	} else if deleteNode.right != nil {
		child = deleteNode.right
	} else {
		child = nil
	}

	if parentNode == nil {
		// 删除的是根节点时
		b.root = child
	} else if parentNode.left == deleteNode {
		parentNode.left = child
	} else {
		parentNode.right = child
	}
}

// 获取值最小节点
func (b *BST) min() *Node {
	node := b.root
	for node.left != nil {
		node = node.left
	}

	return node
}

// 获取值最大节点
func (b *BST) max() *Node {
	node := b.root
	for node.right != nil {
		node = node.right
	}

	return node
}
