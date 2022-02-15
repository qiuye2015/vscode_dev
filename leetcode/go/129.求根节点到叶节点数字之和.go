/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根节点到叶节点数字之和
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
func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, prevSum int) int {
	if root == nil {
		return 0
	}
	prevSum = prevSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return prevSum
	}
	return dfs(root.Left, prevSum) + dfs(root.Right, prevSum)
}

// @lc code=end
