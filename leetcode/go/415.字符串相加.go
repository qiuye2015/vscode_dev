/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

package main

// @lc code=start
func addStrings(num1 string, num2 string) string {
	maxL := max(len(num1), len(num2))
	ret := make([]byte, maxL+1)
	k, i, j := len(ret)-1, len(num1)-1, len(num2)-1
	carry := 0
	for i >= 0 || j >= 0 || carry > 0 {
		if i >= 0 {
			carry += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			carry += int(num2[j] - '0')
			j--
		}
		ret[k] = byte(carry%10 + '0')
		k--
		carry = carry / 10
	}
	if ret[0] == 0 {
		ret = ret[1:]
	}
	return string(ret)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
