/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */
package main

// @lc code=start
func maxProfit(prices []int) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	//dp[i][0]表示第i天交易完后手里没有股票的最大利润
	//dp[i][1]表示第i天交易完后手里持有一支股票的最大利润
	//第0天交易结束的时候dp[0][0]=0，dp[0][1]=−prices[0]
	n := len(prices)
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		tmp := dp0
		dp0 = max(dp0, dp1+prices[i])
		dp1 = max(dp1, tmp-prices[i])
	}
	return dp0
}

// @lc code=end
