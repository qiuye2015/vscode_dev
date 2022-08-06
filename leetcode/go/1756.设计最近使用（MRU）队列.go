/*
 * @lc app=leetcode.cn id=1756 lang=golang
 *
 * [1756] 设计最近使用（MRU）队列
 */

package main

// @lc code=start
type MRUQueue struct {
	nums []int
	tmp  []int
}

func Constructor(n int) MRUQueue {
	vec := make([]int, n)
	for i := 0; i < n; i++ {
		vec[i] = i + 1
	}
	return MRUQueue{
		nums: vec,
	}
}

func (this *MRUQueue) Fetch(k int) int {
	ret := this.nums[k-1]
	this.tmp = this.nums[:k-1]
	this.tmp = append(this.tmp, this.nums[k:]...)
	this.tmp = append(this.tmp, ret)
	this.nums = this.tmp
	return ret
}

/**
 * Your MRUQueue object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Fetch(k);
 */
// @lc code=end
