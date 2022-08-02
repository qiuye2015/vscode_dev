/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

package main

// @lc code=start
func searchRange(nums []int, target int) []int {
	x := lower_bound(nums, target)
	y := lower_bound(nums, target+1)
	if x < len(nums) && nums[x] == target {
		return []int{x, y - 1}
	}
	return []int{-1, -1}
}

//查找大于等于target的位置
func lower_bound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}

// @lc code=end
