/*
 * @lc app=leetcode.cn id=611 lang=golang
 *
 * [611] 有效三角形的个数
 */

package main

import "sort"

// @lc code=start
func triangleNumber(nums []int) int {
	ret := 0
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		left, right := 0, i-1
		for left < right {
			if nums[left]+nums[right] > nums[i] {
				ret += (right - left)
				right--
			} else {
				left++
			}
		}
	}
	return ret
}

// @lc code=end
