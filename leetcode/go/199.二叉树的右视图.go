/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
func rightSideView(root *TreeNode) []int {
	ret := []int{}
	queue := []*TreeNode{root}
	if root == nil {
		return ret
	}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			tmp := queue[i]
			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}
			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
			if i == size-1 {
				ret = append(ret, tmp.Val)
			}
		}
		queue = queue[size:]
	}
	return ret
}

// @lc code=end
