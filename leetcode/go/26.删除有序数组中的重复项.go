/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

package main

// @lc code=start
func removeDuplicates(nums []int) int {
	realSize := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[realSize] = nums[i]
			realSize++
		}
	}
	return realSize
}

// @lc code=end
