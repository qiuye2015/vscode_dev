/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 复制带随机指针的链表
 */

package main

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	dummy := &Node{}
	cur1, cur2 := head, dummy
	for cur1 != nil {
		tmpNode := &Node{Val: cur1.Val}
		m[cur1] = tmpNode
		cur2.Next = tmpNode
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	cur1, cur2 = head, dummy.Next
	for cur1 != nil {
		cur2.Random = m[cur1.Random]
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	return dummy.Next
}

// @lc code=end
