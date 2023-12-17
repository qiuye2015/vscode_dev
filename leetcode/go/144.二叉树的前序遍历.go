/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * 144.二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	// ret := []int{}
	// var dfs func(root *TreeNode)
	// dfs = func(root *TreeNode) {
	// 	if root == nil {
	// 		return
	// 	}
	// 	ret = append(ret, root.Val)
	// 	dfs(root.Left)
	// 	dfs(root.Right)
	// }
	// dfs(root)
	// return ret

	ret := []int{}
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			ret = append(ret, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		last := len(stack) - 1
		node := stack[last]
		stack = stack[:last]
		root = node.Right
	}
	return ret
}

// @lc code=end
