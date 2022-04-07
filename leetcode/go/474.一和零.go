/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */
package main

import "strings"

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, s := range strs {
		zeros := strings.Count(s, "0")
		ones := len(s) - zeros
		for j := m; j >= zeros; j-- {
			for k := n; k >= ones; k-- {
				dp[j][k] = max(dp[j][k], dp[j-zeros][k-ones]+1)
			}
		}
	}
	return dp[m][n]
}

// @lc code=end
