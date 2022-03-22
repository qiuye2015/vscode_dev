/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
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
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	// find mid Node
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// reverse slow List
	var prev *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	l1, l2 := head, prev
	// merge l1 & l2
	// 1->2->nil 4->3->nil
	for l2.Next != nil {
		nextL1 := l1.Next // 2
		nextL2 := l2.Next //3

		l1.Next = l2
		l1 = nextL1

		l2.Next = l1
		l2 = nextL2
	}
}

// @lc code=end
