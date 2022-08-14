/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
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
// func swapPairs(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	next := head.Next
// 	head.Next = swapPairs(head.Next.Next)
// 	next.Next = head
// 	return next
// }
// pre-->1-->2-->3-->4
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for pre != nil && pre.Next != nil && pre.Next.Next != nil {
		n1 := pre.Next
		n2 := pre.Next.Next
		pre.Next = n2     //pre-->2-->3-->4
		n1.Next = n2.Next //1->3->4
		n2.Next = n1      //pre-->2-->1-->3-->4
		pre = n1          //2-->1(pre)-->3-->4
	}
	return dummy.Next
}

// @lc code=end
