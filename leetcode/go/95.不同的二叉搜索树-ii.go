/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
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
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return helper(1, n)
}

// 枚举[start,end]中的值i为当前二叉搜索树的根，
// 那么序列划分为了 [start,i−1] 和 [i+1,end] 两部分
func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	allTrees := []*TreeNode{}
	// 枚举可行根节点
	for i := start; i <= end; i++ {
		leftTrees := helper(start, i-1) // 获得所有可行的左子树集合
		rightTrees := helper(i+1, end)  // 获得所有可行的右子树集合
		// 从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				curTree := &TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				}
				allTrees = append(allTrees, curTree)
			}
		}
	}
	return allTrees
}

// @lc code=end
