/*
 * @lc app=leetcode.cn id=1325 lang=golang
 *
 * [1325] 删除给定值的叶子节点
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
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root != nil {
		//后续遍历; 赋值就是在删除节点
		//把子节点丢到函数里去，如果要删除就返回空，如果不删除就原样返回
		root.Left = removeLeafNodes(root.Left, target)
		root.Right = removeLeafNodes(root.Right, target)
		if root.Left == nil && root.Right == nil && root.Val == target {
			return nil
		}
	}
	return root
}

// @lc code=end
