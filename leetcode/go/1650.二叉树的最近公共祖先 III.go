/*
 * @lc app=leetcode.cn id=1650 lang=golang
 *
 * [1650] 二叉树的最近公共祖先 III
 */

package main

// @lc code=start
/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

func lowestCommonAncestor(p *Node, q *Node) *Node {
	l1, l2 := p, q
	//设两个指针分别从两个给定节点出发，
	//如果两个节点不等，则继续往前一步。
	//如果某个节点到达根节点，则跳到另一个给定的节点。
	//最终两个指针一定会相遇在交点处
	for l1 != l2 {
		if l1.Parent != nil {
			l1 = l1.Parent
		} else {
			l1 = q
		}
		if l2.Parent != nil {
			l2 = l2.Parent
		} else {
			l2 = p
		}
	}
	return l1
}

// @lc code=end
