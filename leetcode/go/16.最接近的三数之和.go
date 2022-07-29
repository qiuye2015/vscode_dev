/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

package main

import "sort"

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	sort.Ints(nums)
	ret := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs(target-sum) < abs(target-ret) {
				ret = sum
			}
			if sum == target {
				return target
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return ret
}

// @lc code=end
