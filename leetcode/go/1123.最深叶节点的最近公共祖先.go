/*
 * @lc app=leetcode.cn id=1123 lang=golang
 *
 * [1123] 最深叶节点的最近公共祖先
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
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var res *TreeNode //
	pre := 0          //保存上一次得到的最近公共祖先的深度

	var dfs func(root *TreeNode, depth int) int
	dfs = func(root *TreeNode, depth int) int {
		if root == nil {
			return depth
		}
		left := dfs(root.Left, depth+1)
		right := dfs(root.Right, depth+1)
		if left == right && left >= pre {
			res = root
			pre = left
		}
		return max(left, right)
	}

	dfs(root, 1)
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
