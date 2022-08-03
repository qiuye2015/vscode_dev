/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */

package main

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	ret := 0
	m, n := len(matrix), len(matrix[0])
	//dp(i,j) 表示以(i,j) 为右下角，且只包含1的正方形的边长最大值
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '0' {
				continue
			}
			if i > 1 && j > 1 && matrix[i-2][j-2] == '1' &&
				matrix[i-2][j-1] == '1' && matrix[i-1][j-2] == '1' {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			} else {
				dp[i][j] = 1
			}
			if dp[i][j] > ret {
				ret = dp[i][j]
			}
		}
	}
	return ret * ret
}

// @lc code=end
