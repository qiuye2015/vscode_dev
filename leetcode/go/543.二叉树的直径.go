/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
func diameterOfBinaryTree(root *TreeNode) int {
	ret := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left) // 左子树最大深度
		right := dfs(root.Right)
		sum := left + right
		ret = max(ret, sum)
		return max(left, right) + 1
	}
	dfs(root)
	return ret
}

// @lc code=end
