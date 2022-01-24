/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */
package main

import "sort"

// @lc code=start
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if nums[i] > 0 { // 因为已经排序好，所以后面不可能有三个数加和等于 00
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue // To prevent the repeat
		}
		target, left, right := -nums[i], i+1, n-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return res
}

// @lc code=end
