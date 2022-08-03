/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

package main

// @lc code=start
func isMatch(s string, p string) bool {
	return helper([]byte(s), []byte(p))
}
func helper(s, p []byte) bool {
	if len(p) == 0 {
		return len(s) == 0
	} else if len(p) > 1 && p[1] == '*' {
		//*匹配0个
		flag1 := helper(s, p[2:])
		//*匹配多个
		flag2 := len(s) > 0 && (p[0] == s[0] || p[0] == '.') && helper(s[1:], p)
		return flag1 || flag2
	} else {
		return len(s) > 0 && (p[0] == s[0] || p[0] == '.') && helper(s[1:], p[1:])
	}
}

// @lc code=end
