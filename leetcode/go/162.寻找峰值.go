/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */

package main

// @lc code=start
func findPeakElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	left, right := 0, n-1
	for {
		mid := left + (right-left)/2
		if mid == 0 {
			if nums[mid] > nums[mid+1] {
				return mid
			}
			left = mid + 1
			continue
		} else if mid == n-1 {
			if nums[mid] > nums[mid-1] {
				return mid
			}
			right = mid - 1
			continue
		}
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		}
		if nums[mid] < nums[mid-1] {
			right = mid - 1
		} else if nums[mid] < nums[mid+1] {
			left = mid + 1
		}
	}
	return -1 //不该到这
}

// @lc code=end
