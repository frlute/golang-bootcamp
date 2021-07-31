package tree

import "fmt"

type BinaryTreeNode struct {
	Data  interface{}
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// TODO 仔细总结下 golang 自定义 String 的用法
func (bt *BinaryTreeNode) String() string {
	return fmt.Sprintf("value:%+v, left:%+v, right:%+v", bt.Data, bt.Left, bt.Right)
}

// TODO 实现 BinaryTreeNode 排序
