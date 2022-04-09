/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */
package main

// @lc code=start
func coinChange(coins []int, amount int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	INF := 100001 //amount + 1
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = INF
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if c <= i {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}

	if dp[amount] == INF {
		return -1
	}
	return dp[amount]
}

// @lc code=end
