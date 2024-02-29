/*
 * @lc app=leetcode.cn id=687 lang=golang
 *
 * 687.最长同值路径
 */

package main

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
	ret := 0

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftSame := dfs(node.Left)
		rightSame := dfs(node.Right)

		left, right := 0, 0
		if node.Left != nil && node.Left.Val == node.Val {
			left = leftSame + 1
		}
		if node.Right != nil && node.Right.Val == node.Val {
			right = rightSame + 1
		}
		ret = max(ret, left+right)
		return max(left, right)
	}
	dfs(root)
	return ret
}

// @lc code=end
