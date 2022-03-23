/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */
package main

// @lc code=start
func isValid(s string) bool {
	st := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' {
			st = append(st, ')')
		} else if c == '[' {
			st = append(st, ']')
		} else if c == '{' {
			st = append(st, '}')
		} else if len(st) == 0 || st[len(st)-1] != c {
			return false
		} else { // st[len(st)-1] == c
			st = st[:len(st)-1]
		}
	}
	return len(st) == 0
}

// @lc code=end
