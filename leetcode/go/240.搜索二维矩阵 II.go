/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */

package main

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix) - 1
	n := len(matrix[0]) - 1
	if target > matrix[m][n] || target < matrix[0][0] {
		return false
	}
	i, j := 0, n //从右上角开始遍历
	for i <= m && j >= 0 {
		if target == matrix[i][j] {
			return true
		} else if target < matrix[i][j] {
			j--
		} else {
			i++
		}
	}
	return false
}

// @lc code=end
