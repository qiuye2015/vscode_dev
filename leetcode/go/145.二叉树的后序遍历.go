/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * 145.二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) []int {
	// ret := make([]int, 0)
	// var dfs func(root *TreeNode)
	// dfs = func(root *TreeNode) {
	// 	if root == nil {
	// 		return
	// 	}
	// 	dfs(root.Left)
	// 	dfs(root.Right)
	// 	ret = append(ret, root.Val)
	// }
	// dfs(root)
	// return ret

	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastVisit *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		last := len(stack) - 1
		node := stack[last]
		if node.Right == nil || node.Right == lastVisit {
			ret = append(ret, node.Val)
			stack = stack[:last]
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return ret
}

// @lc code=end
