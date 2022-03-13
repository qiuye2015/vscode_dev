/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
func isValidBST(root *TreeNode) bool {
	var dfs func(root *TreeNode, minV, maxV int) bool
	dfs = func(root *TreeNode, minV, maxV int) bool {
		if root == nil { // 不会走到
			return true
		}
		if root.Val <= minV || root.Val >= maxV {
			return false
		}
		left := dfs(root.Left, minV, root.Val)
		right := dfs(root.Right, root.Val, maxV)
		return left && right
	}
	return dfs(root, math.MinInt64, math.MaxInt64)
}

// @lc code=end
