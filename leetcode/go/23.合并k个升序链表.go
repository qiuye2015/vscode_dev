/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */
package main

import "container/heap"

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead

	pq := make(PQ, 0)
	for _, node := range lists {
		if node != nil {
			pq = append(pq, node)
		}
	}

	heap.Init(&pq)
	for len(pq) > 0 {
		minNode := heap.Pop(&pq).(*ListNode)
		cur.Next = minNode
		cur = cur.Next
		if minNode.Next != nil {
			heap.Push(&pq, minNode.Next)
		}
	}
	return dummyHead.Next
}

type PQ []*ListNode

func (pq PQ) Len() int {
	return len(pq)
}
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq PQ) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}
func (pq *PQ) Push(x interface{}) {
	node := x.(*ListNode)
	*pq = append(*pq, node)
}
func (pq *PQ) Pop() interface{} {
	ret := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return ret
}

// @lc code=end
