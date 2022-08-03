/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

package main

// @lc code=start
func generate(numRows int) [][]int {
	dp := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		dp[i] = make([]int, i+1)
		dp[i][0], dp[i][i] = 1, 1
		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}
	return dp
}

// @lc code=end
