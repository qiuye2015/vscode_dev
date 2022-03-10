/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */
package main

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func detectCycle(head *ListNode) *ListNode {
	visited := make(map[*ListNode]int, 0)
	for head != nil {
		if _, ok := visited[head]; ok {
			return head
		}
		visited[head] = 1
		head = head.Next
	}
	return nil
}

// @lc code=end
