/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

package main

import (
	"container/heap"
)

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	cnt := map[int]int{}
	for _, num := range nums {
		cnt[num]++
	}
	pq := &PQ{}
	heap.Init(pq)
	for key, val := range cnt {
		heap.Push(pq, item{
			key: key,
			cnt: val,
		})
		if len(*pq) > k {
			heap.Pop(pq)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		x := heap.Pop(pq).(item).key
		ret[k-1-i] = x
	}
	return ret
}

type item struct {
	key, cnt int
}
type PQ []item

func (pq PQ) Len() int {
	return len(pq)
}
func (pq PQ) Less(i, j int) bool {
	return pq[i].cnt < pq[j].cnt
}
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PQ) Push(x interface{}) {
	val := x.(item)
	*pq = append(*pq, val)
}
func (pq *PQ) Pop() interface{} {
	last := len(*pq) - 1
	ret := (*pq)[last]
	*pq = (*pq)[:last]
	return ret
}

// @lc code=end
