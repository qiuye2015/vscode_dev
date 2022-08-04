/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
 */

package main

import "math"

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	return max(robber(root))
}
func robber(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	l1, l2 := robber(root.Left)
	r1, r2 := robber(root.Right)
	selected := root.Val + l2 + r2
	notSelected := max(l1, l2) + max(r1, r2)
	return selected, notSelected
}
func max(values ...int) int {
	maxV := math.MinInt
	for _, v := range values {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

// @lc code=end
