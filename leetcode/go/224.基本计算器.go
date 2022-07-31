/*
 * @lc app=leetcode.cn id=224 lang=golang
 *
 * [224] 基本计算器
 */

package main

// @lc code=start
func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}
	ret, num, sign := 0, 0, 1
	stack := []int{sign}

	for _, char := range s {
		if char == ' ' {
			continue
		} else if char >= '0' && char <= '9' {
			num = num*10 + int(char-'0')
		} else if char == '+' || char == '-' {
			ret += sign * num
			sign = stack[len(stack)-1]
			if char == '-' {
				sign *= -1
			}
			num = 0
		} else if char == '(' { //进入左括号，把左括号之前的符号置于栈顶
			stack = append(stack, sign)
		} else if char == ')' { //退出括号，弹出栈顶符号
			stack = stack[:len(stack)-1]
		}
	}
	ret += sign * num //计算最后一个数
	return ret
}

// @lc code=end
