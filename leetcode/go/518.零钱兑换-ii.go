/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */
package main

// @lc code=start
func change(amount int, coins []int) int {
	// dp[i]表示金额之和等于i的组合数
	dp := make([]int, amount+1)
	dp[0] = 1 //金额之和为0只有不选取任何硬币一种情况
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

// @lc code=end
