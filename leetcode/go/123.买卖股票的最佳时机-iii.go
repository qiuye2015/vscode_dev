/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
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
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}

// @lc code=end
