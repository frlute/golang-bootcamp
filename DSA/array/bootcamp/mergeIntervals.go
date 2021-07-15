package bootcamp

import (
	"fmt"
	"sort"
)

/*
 * [56] Merge Intervals

 解题思路:
	* 区间合并问题
	* 注意临界点
*/
// 暴力法，也就是直接遍历, 时间复杂度 O(n^2) 空间复杂度 O(2n)
func mergeIntervals(intervals [][]int) [][]int {
	length := len(intervals)

	result := make([][]int, 0, length)
	passed := make(map[int]struct{}) // 表示已处理过的 index
	for index, items := range intervals {
		if _, ok := passed[index]; ok {
			continue
		}
		min, max := items[0], items[1]
		for j := index + 1; j < length; j++ {
			if _, ok := passed[j]; ok {
				continue
			}
			l, r := intervals[j][0], intervals[j][1]
			// 表示无交叉的区间
			if l > max || r < min {
				continue
			} else if l >= min && r <= max {
				// 重合，[min, max]
			} else if l <= min && r <= max {
				// 左接壤
				min = l
			} else if l >= min && r > max {
				// 右接壤
				max = r
			} else if l <= min && r >= max {
				min, max = l, r
			}

			// 记录已匹配过的数组
			fmt.Println("----j", j)
			passed[j] = struct{}{}
		}
		fmt.Println("====min, max", min, max)
		result = append(result, []int{min, max})
	}
	fmt.Println(result)
	return result
}

// 上述暴力法存在问题，看完题解后发现自己还是对这种区间问题的理解不够好，导致有些地方没想清楚
/*
首先对二维数组按照左端点升序排序，避免处理过程中区间考虑左端点更小的细节问题, 只处理右端边界即可

时间复杂度:  O(nlog n)，其中 n 为区间的数量。除去排序的开销，我们只需要一次线性扫描，所以主要的时间开销是排序的 O(nlog n)。

空间复杂度: O(logn)，其中 n 为区间的数量。这里计算的是存储答案之外，使用的额外空间。O(log n) 即为排序所需要的空间复杂度。
*/
func mergeIntervalsV1(intervals [][]int) [][]int {
	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}
	prev := intervals[0] // 初始区间

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] < cur[0] { // 无任何重合
			res = append(res, prev)
			prev = cur
		} else {
			// 更新右区间
			prev[1] = max(prev[1], cur[1])
		}
	}
	res = append(res, prev)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
