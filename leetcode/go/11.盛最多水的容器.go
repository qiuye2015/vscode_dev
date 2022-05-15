/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */
package main

// @lc code=start
func maxArea(height []int) int {
	ret, area := 0, 0
	i, j := 0, len(height)-1
	for i < j {
		if height[i] < height[j] {
			area = (j - i) * height[i]
			i++
		} else {
			area = (j - i) * height[j]
			j--
		}
		if area > ret {
			ret = area
		}
	}
	return ret
}

// @lc code=end
