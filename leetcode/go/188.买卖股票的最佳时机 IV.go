/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

package main

// @lc code=start
func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	if k > n/2 {
		return bf(prices)
	}
	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, k+1)
	}
	for i := 1; i <= k; i++ {
		dp[0][i][0] = 0
		dp[0][i][1] = -prices[0]
	}
	for i := 1; i < n; i++ {
		for j := k; j > 0; j-- {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[n-1][k][0]
}

// 不限交易次数k
func bf(prices []int) int {
	ret := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ret += (prices[i] - prices[i-1])
		}
	}
	return ret
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
