/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */
package main

// @lc code=start
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j <= i-1; j++ {
			dp[i] += dp[j] * dp[i-1-j]
		}
	}
	return dp[n]
}

// @lc code=end
