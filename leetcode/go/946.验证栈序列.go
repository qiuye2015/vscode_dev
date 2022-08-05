/*
 * @lc app=leetcode.cn id=946 lang=golang
 *
 * [946] 验证栈序列
 */

package main

// @lc code=start
func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	i, j := 0, 0
	for ; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && j < len(popped) &&
			stack[len(stack)-1] == popped[j] {
			j++
			stack = stack[:len(stack)-1]
		}
	}
	return j == len(popped)
}

// @lc code=end
