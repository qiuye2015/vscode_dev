/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 */

package main

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ans := make([][]int, 0)
	i := 0
	for ; i < len(intervals) && intervals[i][1] < newInterval[0]; i++ {
		ans = append(ans, intervals[i]) //首先将新区间左边且相离的区间加入结果集
	}
	//当前区间是否与新区间重叠 0<=left<=1<=right
	for ; i < len(intervals) && intervals[i][0] <= newInterval[1]; i++ {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
	}
	ans = append(ans, newInterval)
	for i < len(intervals) { //right<0
		ans = append(ans, intervals[i])
		i++
	}
	return ans
}

// @lc code=end
