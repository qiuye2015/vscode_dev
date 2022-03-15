/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
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
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast, slow, pre := head, head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	fast = head
	pre.Next = nil // 前面的断开

	h1 := sortList(fast)
	h2 := sortList(slow)
	ret := mergeList(h1, h2)
	return ret
}

// 合并有序链表
func mergeList(h1, h2 *ListNode) *ListNode {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}
	dummyHead := &ListNode{}
	cur := dummyHead
	for h1 != nil && h2 != nil {
		if h1.Val < h2.Val {
			cur.Next = h1
			h1 = h1.Next
		} else {
			cur.Next = h2
			h2 = h2.Next
		}
		cur = cur.Next
	}
	if h1 != nil {
		cur.Next = h1
	}
	if h2 != nil {
		cur.Next = h2
	}
	return dummyHead.Next
}

// @lc code=end
