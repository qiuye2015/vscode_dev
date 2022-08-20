/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */

package main

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1) // 到达下表i的花费
	dp[0] = 0              // 由于可以选择下标0或1作为初始阶梯
	dp[1] = 0              // 因此有dp[0]=dp[1]=0
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}

// func minCostClimbingStairs(cost []int) int {
// 	n := len(cost)
// 	prev, cur := 0, 0
// 	for i := 2; i <= n; i++ {
// 		next := min(prev+cost[i-2], cur+cost[i-1])
// 		prev, cur = cur, next
// 	}
// 	return cur
// }
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
