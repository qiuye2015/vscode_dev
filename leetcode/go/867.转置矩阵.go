/*
 * @lc app=leetcode.cn id=867 lang=golang
 *
 * [867] 转置矩阵
 */

package main

// @lc code=start
func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ret[j][i] = matrix[i][j]
		}
	}
	return ret
}

// @lc code=end
