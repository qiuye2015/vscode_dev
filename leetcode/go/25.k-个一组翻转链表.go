/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{Next: head}
	pre := dummyHead
	cur := head
	for cur != nil {
		i := 0
		for ; i < k && cur != nil; i++ {
			cur = cur.Next
		}
		if i == k {
			pre.Next = reversePart(head, cur)
			pre = head
			head = cur
		} else {
			pre.Next = head
		}
	}
	return dummyHead.Next
	// cur := head
	// for i := 0; i < k; i++ {
	// 	if cur == nil {
	// 		return head
	// 	}
	// 	cur = cur.Next
	// }
	// newHead := reversePart(head, cur)
	// head.Next = reverseKGroup(cur, k)
	// return newHead
}
func reversePart(start, end *ListNode) *ListNode {
	var pre *ListNode
	cur := start
	for cur != end {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// @lc code=end
