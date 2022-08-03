/*
 * @lc app=leetcode.cn id=1277 lang=golang
 *
 * [1277] 统计全为 1 的正方形子矩阵
 */

package main

// @lc code=start
func countSquares(matrix [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	ret := 0
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1) //dp[i][j]表示以(i,j)为右下角的全1正方形个数
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == 0 {
				continue
			}
			// matrix[i-1][j-1]==1
			if i > 1 && j > 1 && matrix[i-2][j-2] == 1 && matrix[i-2][j-1] == 1 && matrix[i-1][j-2] == 1 {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			} else {
				dp[i][j] = 1
			}
			// if i == 1 || j == 1 {
			// 	dp[i][j] = 1
			// } else if matrix[i-2][j-2] == 1 && matrix[i-2][j-1] == 1 && matrix[i-1][j-2] == 1 {
			// 	dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			// } else {
			// 	dp[i][j] = 1
			// }
			ret += dp[i][j]
		}
	}
	return ret
}

// @lc code=end
