/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

package main

// @lc code=start
func firstMissingPositive(nums []int) int {
	for i := range nums {
		for nums[i] >= 1 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			// 把nums[i] 放在 nums[i]-1 所在的位置上（eg:把1放在0上）
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

// @lc code=end
