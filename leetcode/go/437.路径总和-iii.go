/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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
//  时间复杂度：O(N^2)
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	ret := rootSum(root, targetSum)      //选择当前节点
	ret += pathSum(root.Left, targetSum) //不选当前节点
	ret += pathSum(root.Right, targetSum)
	return ret
}

// 表示以节点p为起点向下且满足路径总和为val的路径数目
func rootSum(p *TreeNode, val int) int {
	ret := 0
	if p == nil {
		return ret
	}
	if p.Val == val {
		ret++
	}
	ret += rootSum(p.Left, val-p.Val)
	ret += rootSum(p.Right, val-p.Val)
	return ret
}

// @lc code=end
