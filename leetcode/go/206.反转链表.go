/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
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
func reverseList(head *ListNode) *ListNode {
	var helper func(head, pre *ListNode) *ListNode
	helper = func(head, pre *ListNode) *ListNode {
		if head == nil {
			return pre
		}
		next := head.Next
		head.Next = pre
		pre = head
		return helper(next, pre)
	}
	return helper(head, nil)
}

// @lc code=end
