/*
 * @lc app=leetcode.cn id=905 lang=golang
 *
 * [905] 按奇偶排序数组
 */

package main

// @lc code=start
func sortArrayByParity(nums []int) []int {
	start := 0
	end := len(nums) - 1

	for start < end {
		if nums[start]%2 != 0 { // odd; swap with end
			nums[start], nums[end] = nums[end], nums[start]
			end--
		} else {
			start++
		}
	}
	return nums
}

// @lc code=end
