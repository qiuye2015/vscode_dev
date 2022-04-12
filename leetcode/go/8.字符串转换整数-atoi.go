/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */
package main

// @lc code=start
func myAtoi(s string) int {
	//空格，非数字，正负号
	idx, n := 0, len(s)
	ret, sign := 0, 1
	maxInt := 1<<31 - 1 //最大正数 math.MaxInt32
	minInt := -1 << 31  //最小负数 math.MinInt32
	//空格
	for idx < n && (s[idx] == ' ') {
		idx++
	}
	if idx == n {
		return 0
	}
	//正负号
	if s[idx] == '+' {
		sign = 1
		idx++
	} else if s[idx] == '-' {
		sign = -1
		idx++
	}
	for idx < n && s[idx] >= '0' && s[idx] <= '9' {
		ret = ret*10 + int(s[idx]-'0')
		idx++
		if sign*ret > maxInt {
			return maxInt
		}
		if sign*ret < minInt {
			return minInt
		}
	}
	return ret * sign
}

// 42
// 0042
// +-12
// -91283472332
// 2147483647
// -2147483648

// @lc code=end
