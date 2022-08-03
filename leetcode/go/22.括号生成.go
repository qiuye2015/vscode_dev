/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

package main

// @lc code=start
func generateParenthesis(n int) []string {
	ret := []string{}
	//left 左括号还有几个可以使用
	var helper func(left, right int, cur string)
	helper = func(left, right int, cur string) {
		if left == 0 && right == 0 {
			ret = append(ret, string(cur))
			return
		}
		//此时生成的字符串中右括号的个数大于左括号的个数，即会出现')('这样的非法串
		if left > right {
			return
		}
		if left > 0 {
			helper(left-1, right, cur+"(")
		}
		if right > 0 {
			helper(left, right-1, cur+")")
		}
	}
	helper(n, n, "")
	return ret
}

// @lc code=end
