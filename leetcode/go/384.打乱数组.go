/*
 * @lc app=leetcode.cn id=384 lang=golang
 *
 * [384] 打乱数组
 */

package main

import (
	"math/rand"
	"time"
)

// @lc code=start
type Solution struct {
	original []int
}

func Constructor(nums []int) Solution {
	rand.Seed(time.Now().UnixNano())
	return Solution{
		original: nums,
	}
}
func (this *Solution) Reset() []int {
	return this.original
}
func (this *Solution) Shuffle() []int {
	n := len(this.original)
	ret := make([]int, n)
	copy(ret, this.original)
	for i := 0; i < n; i++ {
		j := i + rand.Intn(n-i)
		ret[i], ret[j] = ret[j], ret[i]
	}
	return ret
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
// @lc code=end
