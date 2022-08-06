/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */

package main

import (
	"strconv"
)

// @lc code=start
func restoreIpAddresses(s string) []string {
	ret := []string{}
	if len(s) > 12 {
		return ret
	}
	var dfs func(s, path string, blocks int)
	dfs = func(s, path string, blocks int) {
		if blocks > 4 {
			return
		}
		if blocks == 4 && len(s) == 0 {
			ret = append(ret, path[:len(path)-1]) //去除最后一个.
			return
		}
		for i := 1; i < 4; i++ {
			if i > len(s) {
				return
			}
			num, _ := strconv.Atoi(s[:i])
			if s[:i] == "0" || (s[0] != '0' && num < 256) {
				dfs(s[i:], path+s[:i]+".", blocks+1)
			}
		}
	}
	dfs(s, "", 0)
	return ret
}

// @lc code=end
