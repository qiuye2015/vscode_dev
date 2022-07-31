/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

package main

// @lc code=start
func longestValidParentheses(s string) int {
	ret := 0
	st := []int{-1}
	for idx, char := range s {
		if char == '(' {
			st = append(st, idx) //push
		} else {
			st = st[:len(st)-1] //pop
			if len(st) == 0 {
				st = append(st, idx)
			} else {
				newLen := idx - st[len(st)-1]
				if newLen > ret {
					ret = newLen
				}
			}
		}
	}
	return ret
}

// @lc code=end
