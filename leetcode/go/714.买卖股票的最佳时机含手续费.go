/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */

package main

// @lc code=start
func maxProfit(prices []int, fee int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0], dp[0][1] = 0, -prices[0]-fee
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]-fee)
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
