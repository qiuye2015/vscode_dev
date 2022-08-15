/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */

package main

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	ret := []int{}
	for i := range nums {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := range nums {
		if nums[i] != i+1 {
			ret = append(ret, i+1)
		}
	}
	return ret
}

// @lc code=end
