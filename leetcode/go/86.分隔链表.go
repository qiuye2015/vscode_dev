/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
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
func partition(head *ListNode, x int) *ListNode {
	dummy1, dummy2 := &ListNode{}, &ListNode{}
	l1, l2 := dummy1, dummy2
	cur := head
	for cur != nil {
		if cur.Val < x {
			l1.Next = cur
			l1 = l1.Next
		} else {
			l2.Next = cur
			l2 = l2.Next
		}
		cur = cur.Next
	}
	l1.Next = dummy2.Next
	l2.Next = nil
	return dummy1.Next
}

// @lc code=end
