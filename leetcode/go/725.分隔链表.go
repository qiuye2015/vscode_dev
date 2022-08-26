/*
 * @lc app=leetcode.cn id=725 lang=golang
 *
 * [725] 分隔链表
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
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
	ret := make([]*ListNode, k)
	count := listLen(head)
	x, y := count/k, count%k
	idx := 0
	for head != nil {
		cnt := x
		if y > 0 {
			cnt += 1
			y -= 1
		}
		ret[idx] = head
		idx++
		for head != nil && cnt > 1 {
			head = head.Next
			cnt--
		}
		if head != nil {
			next := head.Next
			head.Next = nil
			head = next
		}
	}
	return ret
}
func listLen(head *ListNode) int {
	ret := 0
	for head != nil {
		ret++
		head = head.Next
	}
	return ret
}

// @lc code=end
