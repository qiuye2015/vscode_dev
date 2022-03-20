/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	var dfs func(root *TreeNode, target int, cur []int)
	dfs = func(root *TreeNode, target int, cur []int) {
		if root == nil {
			return
		}
		cur = append(cur, root.Val)
		target -= root.Val
		if root.Left == nil && root.Right == nil &&
			target == 0 {
			tmp := append([]int{}, cur...)
			res = append(res, tmp)
			return
		}
		dfs(root.Left, target, cur)
		dfs(root.Right, target, cur)
	}
	dfs(root, targetSum, []int{})
	return res
}

// @lc code=end
