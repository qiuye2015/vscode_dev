/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */
package main

// @lc code=start
func trap(height []int) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	n, ret := len(height), 0
	maxFromLeft := make([]int, n) //表示下标i左边的位置中height的最大高度
	maxFromRight := make([]int, n)
	for i := 1; i < n; i++ {
		maxFromLeft[i] = max(maxFromLeft[i-1], height[i-1])
	}
	for i := n - 2; i >= 0; i-- {
		maxFromRight[i] = max(maxFromRight[i+1], height[i+1])
	}
	for i := 1; i < n-1; i++ {
		minHeight := min(maxFromLeft[i], maxFromRight[i])
		if minHeight > height[i] {
			ret += minHeight - height[i]
		}
	}
	return ret
}

// @lc code=end
