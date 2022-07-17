/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

package main

// @lc code=start
func findKthLargest(nums []int, k int) int {
	start, end := 0, len(nums)-1
	for {
		index := partition(nums, start, end)
		if index == k-1 {
			return nums[index]
		} else if index < k {
			start = index + 1
		} else {
			end = index - 1
		}
	}
	return -1
}

// 从大到小排序
func partition(nums []int, start, end int) int {
	pivot := nums[start]
	for start < end {
		for start < end && nums[end] <= pivot {
			end--
		}
		nums[start] = nums[end]
		for start < end && nums[start] >= pivot {
			start++
		}
		nums[end] = nums[start]
	}
	nums[start] = pivot
	return start
}

// @lc code=end
