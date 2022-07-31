/*
 * @lc app=leetcode.cn id=581 lang=golang
 *
 * [581] 最短无序连续子数组
 */

package main

import (
	"sort"
)

// @lc code=start
func findUnsortedSubarray(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	left, right := 0, len(nums)-1
	tmp := append([]int{}, nums...)
	sort.Ints(tmp)
	for left <= right && tmp[left] == nums[left] {
		left++
	}
	for left <= right && tmp[right] == nums[right] {
		right--
	}
	return right - left + 1
}

// @lc code=end
