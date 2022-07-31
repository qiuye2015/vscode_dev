/*
 * @lc app=leetcode.cn id=945 lang=golang
 *
 * [945] 使数组唯一的最小增量
 */

package main

import "sort"

// @lc code=start
func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	res := 0
	preNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == preNum+1 {
			//preNum + 1 表示当前数「最好」是这个值
			preNum = nums[i]
		} else if nums[i] > preNum+1 {
			//当前这个数已经足够大，这种情况可以合并到上一个分支
			preNum = nums[i]
		} else { //  nums[i] < preNum+1
			res += preNum + 1 - nums[i]
			preNum++
		}
	}
	return res
}

// @lc code=end
