package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_evenOdd(t *testing.T) {
	// TODO 规范测试用例
	assert.Equal(t, []int{2, 4, 8, 3, 9, 5, 1}, evenOdd([]int{1, 5, 8, 3, 4, 9, 2}))
}

type inputArgs struct {
	A []int
	B []int
}

func Test_intersect(t *testing.T) {
	testCases := []struct {
		args   inputArgs
		expect []int
	}{
		{
			args: inputArgs{
				A: []int{2, 3, 3, 5, 7, 11},
				B: []int{3, 3, 7, 15, 31},
			},
			expect: []int{3, 7},
		},
	}

	for _, c := range testCases {
		actual := intersect(c.args.A, c.args.B)
		assert.Equal(t, c.expect, actual)
	}
}

func Test_twoSum(t *testing.T) {
	testCases := []struct {
		input  []int
		target int
		output []int
	}{
		{
			input:  []int{1},
			target: 5,
			output: nil,
		},
		{
			input:  []int{1, 2, 5, 6, 8, 10},
			target: 13,
			output: []int{2, 4},
		},
	}

	for _, item := range testCases {
		t.Run("test two sum", func(t *testing.T) {
			actual := twoSum(item.input, item.target)
			assert.Equal(t, item.output, actual)

			assert.Equal(t, item.output, twoSumBasedHash(item.input, item.target))

			// 此解法数组必须已排序
			assert.Equal(t, item.output, twoSumBasedInvariants(item.input, item.target))
		})
	}
}

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
