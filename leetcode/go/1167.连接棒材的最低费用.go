/*
 * @lc app=leetcode.cn id=1167 lang=golang
 *
 * [1167] 连接棒材的最低费用
 */

package main

import "container/heap"

// @lc code=start
func connectSticks(sticks []int) int {
	ret := 0
	if len(sticks) < 2 {
		return ret
	}
	pq := new(MinHeap)
	for _, v := range sticks {
		heap.Push(pq, v)
	}

	for pq.Len() > 1 {
		x := heap.Pop(pq).(int)
		y := heap.Pop(pq).(int)
		heap.Push(pq, x+y)
		ret += x + y
	}
	return ret
}

type MinHeap []int

func (pq MinHeap) Len() int {
	return len(pq)
}
func (pq MinHeap) Less(i, j int) bool {
	return pq[i] < pq[j]
}
func (pq MinHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *MinHeap) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}
func (pq *MinHeap) Pop() interface{} {
	last := len(*pq) - 1
	ret := (*pq)[last]
	(*pq) = (*pq)[:last]
	return ret
}

// @lc code=end
