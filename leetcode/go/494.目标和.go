/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */
package main

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	//nums求和：sum
	//nums负数求和：neg
	//nums正数求和：sum-neg
	//sum-neg -neg =target==> sum-2*neg = target
	sum := 0
	for _, v := range nums {
		sum += v
	}
	diff := sum - target //必须大于等于0&为偶数
	if diff < 0 || diff&0x01 == 1 {
		return 0
	}
	neg := diff / 2

	dp := make([]int, neg+1)
	dp[0] = 1
	for _, num := range nums {
		for j := neg; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[neg]
}

// @lc code=end
