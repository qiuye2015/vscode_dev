/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

package main

import (
	"strings"
)

// @lc code=start
func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	res := make([]string, numRows)
	row, flag := 0, -1
	for i := range s {
		res[row] = res[row] + string(s[i])
		if row == 0 || row == numRows-1 {
			flag = -flag
		}
		row += flag
	}
	return strings.Join(res, "")
}

// @lc code=end
