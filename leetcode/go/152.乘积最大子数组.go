/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */
package main

// @lc code=start
func maxProduct(nums []int) int {
	ret, maxV, minV := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			maxV = maxInt(nums[i], nums[i]*maxV)
			minV = minInt(nums[i], nums[i]*minV)
		} else if nums[i] == 0 {
			maxV = 0
			minV = 0
		} else {
			tmp := maxV
			maxV = maxInt(nums[i], nums[i]*minV)
			minV = minInt(nums[i], nums[i]*tmp)
		}
		ret = maxInt(maxV, ret)
	}
	return ret
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
