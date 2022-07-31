/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

package main

// @lc code=start
type MinStack struct {
	nums []int
	min  []int
}

func Constructor() MinStack {
	return MinStack{
		nums: make([]int, 0),
		min:  make([]int, 0),
	}
}
func (this *MinStack) Push(val int) {
	this.nums = append(this.nums, val)
	if len(this.min) == 0 || this.GetMin() > val {
		this.min = append(this.min, val)
	} else {
		this.min = append(this.min, this.GetMin())
	}
}
func (this *MinStack) Pop() {
	last := len(this.nums) - 1
	this.nums = this.nums[:last]

	last = len(this.min) - 1
	this.min = this.min[:last]
}
func (this *MinStack) Top() int {
	last := len(this.nums) - 1
	return this.nums[last]
}
func (this *MinStack) GetMin() int {
	last := len(this.min) - 1
	return this.min[last]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
