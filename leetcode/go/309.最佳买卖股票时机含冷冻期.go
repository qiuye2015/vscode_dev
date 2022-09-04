/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

package main

// @lc code=start
func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0], dp[0][1] = 0, -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		if i == 1 {
			dp[i][1] = max(dp[i-1][1], -prices[i])
		} else {
			dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
		}
	}
	return dp[n-1][0]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
