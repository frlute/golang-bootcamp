package bootcamp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeElement(t *testing.T) {
	type Case struct {
		name   string
		input  []int
		target int // 指定删除的元素
		output int // 删除元素后的数组长度
	}

	cases := []Case{
		{
			name:   "normal remove element",
			input:  []int{1, 2, 3, 5, 3, 2},
			target: 3,
			output: 4,
		},
		{
			name:   "删除元素在首尾",
			input:  []int{11, 12, 13, 15, 13, 11},
			target: 11,
			output: 4,
		},
		{
			name:   "删除元素在尾部",
			input:  []int{21, 22, 23, 25, 23, 24},
			target: 24,
			output: 5,
		},
		{
			name:   "error case",
			input:  []int{1, 2, 1},
			target: 2,
			output: 0,
		},
	}

	type removeFunc func([]int, int) int
	removeFunctions := []removeFunc{
		removeElement,
		removeElementV1,
		removeElementV2,
	}

	for _, fn := range removeFunctions {
		for _, item := range cases {
			// 因原地修改数组，故需要每次调用时把参数隔离开，避免不同删除删除相互影响
			input := make([]int, len(item.input))
			copy(input, item.input)

			t.Run(item.name, func(t *testing.T) {
				res := fn(input, item.target)

				if item.name != "error case" {
					assert.Equal(t, item.output, res)
				}
				// 验证下数组内容是否符合要求,避免实现只是计元素个数, 如在 removeElement 中注释赋值的语句即可出现此错误
				if item.name != "删除元素在尾部" {
					assert.NotEqual(t, item.input, input)
				}
			})
		}
	}
}
