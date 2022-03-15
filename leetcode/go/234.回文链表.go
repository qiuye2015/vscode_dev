/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
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
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	fast, slow := head, head
	// 1,2   1,2,1
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// skip central node if length of linked list is odd number
	if fast != nil { //可有可无
		slow = slow.Next
	}
	newHead := reverseList2(slow)

	for newHead != nil {
		if head.Val != newHead.Val {
			return false
		}
		head = head.Next
		newHead = newHead.Next
	}
	return true
}
func reverseList2(head *ListNode) *ListNode {
	// 1--> 2--> 3
	if head == nil {
		return nil
	}
	var pre, next *ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// @lc code=end
