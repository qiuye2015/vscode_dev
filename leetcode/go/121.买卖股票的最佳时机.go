/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */
package main

// @lc code=start
func maxProfit(prices []int) int {
	minP := prices[0]
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minP {
			minP = prices[i]
		}
		profile := prices[i] - minP
		if profile > res {
			res = profile
		}
	}
	return res
}

// @lc code=end
