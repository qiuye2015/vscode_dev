/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

package main

// @lc code=start
func sortArray(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	QuickSort(nums, 0, len(nums)-1)
	return nums
}

func QuickSort(nums []int, start, end int) {
	if start < end {
		point := partition(nums, start, end)
		QuickSort(nums, start, point-1)
		QuickSort(nums, point+1, end)
	}
}
func partition(nums []int, start, end int) int {
	point := nums[start]
	for start < end {
		for start < end && nums[end] > point {
			end--
		}
		nums[start], nums[end] = nums[end], nums[start]
		for start < end && nums[start] < point {
			start++
		}
		nums[start], nums[end] = nums[end], nums[start]
	}
	nums[start] = point
	return start
}

// @lc code=end
