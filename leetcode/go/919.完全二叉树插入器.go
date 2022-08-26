/*
 * @lc app=leetcode.cn id=919 lang=golang
 *
 * [919] 完全二叉树插入器
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
type CBTInserter struct {
	root      *TreeNode
	lastLevel []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	level := []*TreeNode{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left == nil || node.Right == nil {
			level = append(level, node)
		}
	}
	return CBTInserter{root: root, lastLevel: level}
}

func (this *CBTInserter) Insert(val int) int {
	tmp := &TreeNode{Val: val}
	node := this.lastLevel[0]
	if node.Left == nil {
		node.Left = tmp
	} else {
		node.Right = tmp
		this.lastLevel = this.lastLevel[1:]
	}
	this.lastLevel = append(this.lastLevel, tmp)
	return node.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
// @lc code=end
