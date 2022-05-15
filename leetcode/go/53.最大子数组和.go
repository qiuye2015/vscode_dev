/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */
package main

// @lc code=start
func maxSubArray(nums []int) int {
	max, sum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

// @lc code=end
