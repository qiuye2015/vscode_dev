/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := GetListLen(headA)
	lenB := GetListLen(headB)
	if lenA < lenB { //保证len A>B
		headA, headB = headB, headA
		lenA, lenB = lenB, lenA
	}
	diff := lenA - lenB
	for i := 0; i < diff; i++ {
		headA = headA.Next
	}
	for headA != nil && headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	if headA != nil {
		return headA
	}
	return nil
}
func GetListLen(head *ListNode) int {
	count := 0
	for head != nil {
		count++
		head = head.Next
	}
	return count
}

// @lc code=end
