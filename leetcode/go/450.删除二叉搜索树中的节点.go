/*
 * @lc app=leetcode.cn id=450 lang=golang
 *
 * [450] 删除二叉搜索树中的节点
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
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key == root.Val {
		if root.Left == nil { //无左子：其右子顶替其位置，删除了该节点
			return root.Right
		} else if root.Right == nil { //无右子：其左子顶替其位置，删除了该节点
			return root.Left
		}
		//左右子节点都有
		minNode := root.Right
		for minNode.Left != nil { //寻找欲删除节点右子树的最左节点
			minNode = minNode.Left
		}
		//minNode.Left = root.Left // 将欲删除节点的左子树成为其右子树的最左节点的左子树
		//root = root.Right        // 欲删除节点的右子顶替其位置，节点被删除
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, minNode.Val)
	}
	return root
}

// @lc code=end
