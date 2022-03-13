/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] 恢复二叉搜索树
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
func recoverTree(root *TreeNode) {
	var frist, second, pre *TreeNode

	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		if pre != nil && root.Val < pre.Val {
			if frist == nil {
				frist = pre
			}
			second = root
		}
		pre = root
		inorder(root.Right)
	}
	inorder(root)
	frist.Val, second.Val = second.Val, frist.Val
}

// @lc code=end
