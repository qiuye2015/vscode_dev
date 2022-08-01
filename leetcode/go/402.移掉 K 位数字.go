/*
 * @lc app=leetcode.cn id=402 lang=golang
 *
 * [402] 移掉 K 位数字
 */

package main

import "strings"

// @lc code=start
func removeKdigits(num string, k int) string {
	stack := make([]rune, 0)
	for _, ch := range num {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > ch {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, ch)
	}
	// if k > 0 {
	stack = stack[:len(stack)-k]
	// }
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		return "0"
	}
	return ans
}

// @lc code=end
