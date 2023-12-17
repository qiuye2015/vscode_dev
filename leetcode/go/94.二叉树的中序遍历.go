/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * 94.二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	// ret := make([]int, 0)
	// var dfs func(root *TreeNode)
	// dfs = func(root *TreeNode) {
	// 	if root == nil {
	// 		return
	// 	}
	// 	dfs(root.Left)
	// 	ret = append(ret, root.Val)
	// 	dfs(root.Right)
	// }
	// dfs(root)
	// return ret

	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		last := len(stack) - 1
		node := stack[last]
		stack = stack[:last]
		ret = append(ret, node.Val)
		root = node.Right
	}
	return ret
}

// @lc code=end
