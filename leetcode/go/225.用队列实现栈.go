/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */

package main

// @lc code=start
type MyStack struct {
	queue1 Queue // 主要 后进的放入前面
	queue2 Queue // 辅助
}

func Constructor() MyStack {
	return MyStack{
		queue1: Queue{[]int{}},
		queue2: Queue{[]int{}},
	}
}

func (this *MyStack) Push(x int) {
	this.queue2.push(x)
	for this.queue1.size() > 0 {
		x = this.queue1.pop()
		this.queue2.push(x)
	}
	this.queue1, this.queue2 = this.queue2, this.queue1
}

func (this *MyStack) Pop() int {
	return this.queue1.pop()
}

func (this *MyStack) Top() int {
	return this.queue1.peek()
}

func (this *MyStack) Empty() bool {
	return this.queue1.empty()
}

// 队列;先进先出
type Queue struct {
	Data []int
}

func (q *Queue) push(x int) {
	q.Data = append(q.Data, x)
}
func (q *Queue) pop() int {
	if len(q.Data) > 0 {
		x := q.Data[0]
		q.Data = q.Data[1:]
		return x
	}
	return 0 //err
}
func (q *Queue) peek() int {
	if len(q.Data) > 0 {
		return q.Data[0]
	}
	return 0 //err
}
func (q *Queue) size() int {
	return len(q.Data)
}
func (q *Queue) empty() bool {
	return len(q.Data) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end
