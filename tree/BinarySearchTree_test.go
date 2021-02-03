package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBST(t *testing.T) {
	bst := NewBST(1)
	ast := assert.New(t)

	bst.insert(7)
	bst.insert(2)

	ast.Equal(2, bst.find(2).data.(int))
	ast.Equal(7, bst.max().data.(int))
	ast.Equal(1, bst.min().data.(int))

	// TODO 注意这儿判断 nil 的用法，不能直接用 ast.Equal(nil, bst.find(7))
	ast.Equal(true, bst.find(9) == nil)
	ast.Nil(bst.find(9))

	// 删除没有子节点的节点
	bst.delete(2)
	ast.Nil(bst.find(2))

	// 删除只有一个节点的节点
	bst.insert(2)
	bst.insert(5)

	ast.Equal(7, bst.find(7).data.(int))
	bst.delete(7)

	ast.Nil(bst.find(7))

	// 删除有两个节点的节点
	bst.insert(7)
	bst.insert(6)
	bst.insert(5)

	ast.Equal(7, bst.find(7).data.(int))
	ast.Equal(6, bst.find(6).data.(int))
	ast.Equal(5, bst.find(5).data.(int))

	bst.delete(7)
	ast.Nil(bst.find(7))
	ast.Equal(6, bst.max().data.(int))

}
