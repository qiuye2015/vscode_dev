/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */

package main

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	left := 0  //剩余油量
	start := 0 //初始选索引0作为起点
	totalGas, totalCost := 0, 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		left += gas[i] - cost[i]
		if left < 0 { //去不了下一站，0到i都不是起点
			start = i + 1
			left = 0
		}
	}
	if totalCost > totalGas { //总油量不够，问题无解
		return -1
	}
	return start
}

// @lc code=end
