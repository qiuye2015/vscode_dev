/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

package main

import "sort"

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		threeSum_(nums[i+1:], target-nums[i], nums[i], &res)
	}
	return res
}

func threeSum_(nums []int, target int, first int, res *[][]int) {
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				*res = append(*res, []int{first, nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < target {
				left++
			} else if sum > target {
				right--
			}
		}
	}
}

// @lc code=end
