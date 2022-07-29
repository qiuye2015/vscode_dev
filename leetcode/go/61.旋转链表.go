/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
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
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	slow, fast := head, head
	for i := 0; i < k; i++ {
		if fast.Next == nil { //i+1为链表的长度
			return rotateRight(head, k%(i+1))
		}
		fast = fast.Next
	}
	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}
	newHead := slow.Next
	slow.Next, fast.Next = nil, head
	return newHead
}

// @lc code=end
