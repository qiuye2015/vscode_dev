/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * 429.N 叉树的层序遍历
 */

package main

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	ret := make([][]int, 0)
	stack := []*Node{root}
	for len(stack) > 0 {
		level := []int{}
		length := len(stack)
		for i := 0; i < length; i++ {
			node := stack[i]
			level = append(level, node.Val)
			if len(node.Children) > 0 {
				stack = append(stack, node.Children...)
			}
		}
		stack = stack[length:]
		ret = append(ret, level)
	}
	return ret
}

// @lc code=end
