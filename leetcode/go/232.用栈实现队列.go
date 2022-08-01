/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

package main

// @lc code=start
type MyQueue struct {
	stackPush, stackPop []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.stackPush = append(this.stackPush, x)
}

func (this *MyQueue) Pop() int {
	ans := this.Peek()
	this.stackPop = this.stackPop[:len(this.stackPop)-1]
	return ans
}

func (this *MyQueue) Peek() int {
	if len(this.stackPop) == 0 {
		for len(this.stackPush) != 0 {
			last := len(this.stackPush) - 1
			top := this.stackPush[last]
			this.stackPush = this.stackPush[:last]
			this.stackPop = append(this.stackPop, top)
		}
	}
	return this.stackPop[len(this.stackPop)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stackPop) == 0 && len(this.stackPush) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end
