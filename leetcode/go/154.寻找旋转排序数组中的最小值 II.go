/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */

package main

// @lc code=start
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] { //右边无序
			left = mid + 1
		} else if nums[mid] < nums[right] { //左边无序
			right = mid
		} else if nums[mid] == nums[right] { //diff
			right = right - 1
		}
	}
	return nums[left]
}

// @lc code=end
