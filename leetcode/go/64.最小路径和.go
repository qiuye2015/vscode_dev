/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

package main

// @lc code=start
func minPathSum(grid [][]int) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	m := len(grid)
	n := len(grid[0])

	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}

// @lc code=end
