/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * 590.N 叉树的后序遍历
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

func postorder(root *Node) []int {
	// ret := make([]int, 0)
	// var dfs func(root *Node)
	// dfs = func(root *Node) {
	// 	if root == nil {
	// 		return
	// 	}
	// 	for _, ch := range root.Children {
	// 		dfs(ch)

	// 	}
	// 	ret = append(ret, root.Val)
	// }
	// dfs(root)
	// return ret
	if root == nil {
		return nil
	}
	ret := make([]int, 0)
	stack := []*Node{root}
	visited := map[*Node]bool{}
	for len(stack) > 0 {
		last := len(stack) - 1
		node := stack[last]
		//当前节点为叶子节点或者已经访问过了
		if len(node.Children) == 0 || visited[node] {
			stack = stack[:last]
			ret = append(ret, node.Val)
		} else {
			for i := len(node.Children) - 1; i >= 0; i-- {
				stack = append(stack, node.Children[i])
			}
			visited[node] = true
		}
	}
	return ret
}

// @lc code=end
