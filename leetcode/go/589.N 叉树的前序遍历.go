/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * 589.N 叉树的前序遍历
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

func preorder(root *Node) []int {
	// ret := make([]int, 0)
	// var dfs func(root *Node)
	// dfs = func(root *Node) {
	// 	if root == nil {
	// 		return
	// 	}
	// 	ret = append(ret, root.Val)
	// 	for _, node := range root.Children {
	// 		dfs(node)
	// 	}
	// }
	// dfs(root)
	// return ret

	if root == nil {
		return nil
	}
	ret := make([]int, 0)
	stack := []*Node{root}
	for len(stack) > 0 {
		last := len(stack) - 1
		node := stack[last]
		stack = stack[:last]
		ret = append(ret, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return ret
}

// @lc code=end
