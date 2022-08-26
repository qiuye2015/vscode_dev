/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
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
func countNodes(root *TreeNode) int {
	if root == nil { //O(logn * logn)
		return 0
	}
	leftHeigh := getHeight(root.Left)
	rightHeigh := getHeight(root.Right)
	res := 1
	//左子树一定是满二叉树，因为节点已经填充到右子树了，左子树必定已经填满了。
	//所以左子树的节点总数我们可以直接得到，是 2^left - 1，
	//加上当前这个root节点，则正好是 2^left
	if leftHeigh == rightHeigh {
		res += (1 << leftHeigh) - 1
		res += countNodes(root.Right)
	} else {
		//此时最后一层不满，但倒数第二层已经满了，可以直接得到右子树的节点个数
		res += (1 << rightHeigh) - 1
		res += countNodes(root.Left)
	}
	return res
}

// 统计树的深度
func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + getHeight(node.Left)
}

// @lc code=end
