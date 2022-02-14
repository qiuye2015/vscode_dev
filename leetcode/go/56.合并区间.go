/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */
package main

import "sort"

// @lc code=start
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ret := make([][]int, 0)
	ret = append(ret, intervals[0])
	for i := 1; i < len(intervals); i++ {
		length := len(ret)
		last := ret[length-1]
		if intervals[i][0] > last[1] { // 新区间左端点 > 旧区间右端点时，无交集
			ret = append(ret, intervals[i])
		} else if intervals[i][1] > last[1] { // 取两个区间右端点的最大值
			ret[length-1][1] = intervals[i][1]
		}
	}
	return ret
}

// @lc code=end
