/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

package main

import (
	"strings"
)

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var sb strings.Builder
	for i := 0; ; i++ {
		for _, s := range strs {
			if i == len(s) || strs[0][i] != s[i] {
				return sb.String()
			}
		}
		sb.WriteByte(strs[0][i])
	}
}

// @lc code=end
