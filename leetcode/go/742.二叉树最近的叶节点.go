/*
 * @lc app=leetcode.cn id=742 lang=golang
 *
 * [742] 二叉树最近的叶节点
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
func findClosestLeaf(root *TreeNode, k int) int {
	var target *TreeNode
	parents := map[*TreeNode]*TreeNode{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if target == nil && root != nil && root.Val == k {
			target = root
		}
		if root.Left != nil {
			parents[root.Left] = root
			dfs(root.Left)
		}
		if root.Right != nil {
			parents[root.Right] = root
			dfs(root.Right)
		}
	}
	bfs := func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		res := -1
		queue := []*TreeNode{}
		seen := map[*TreeNode]int{}
		queue = append(queue, root)
		seen[root] = 1
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left == nil && node.Right == nil {
				res = node.Val
				break
			}
			if node.Left != nil && seen[node.Left] == 0 {
				seen[node.Left] = 1
				queue = append(queue, node.Left)
			}
			if node.Right != nil && seen[node.Right] == 0 {
				seen[node.Right] = 1
				queue = append(queue, node.Right)
			}
			if parents[node] != nil && seen[parents[node]] == 0 {
				seen[parents[node]] = 1
				queue = append(queue, parents[node])
			}
		}
		return res
	}
	dfs(root)
	return bfs(target)
}

// @lc code=end
