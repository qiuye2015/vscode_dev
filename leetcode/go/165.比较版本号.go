/*
 * @lc app=leetcode.cn id=165 lang=golang
 *
 * [165] 比较版本号
 */

package main

import (
	"strconv"
	"strings"
)

// @lc code=start
func compareVersion(version1 string, version2 string) int {
	vec1, vec2 := strings.Split(version1, "."), strings.Split(version2, ".")
	for i, j := 0, 0; i < len(vec1) || j < len(vec2); i, j = i+1, j+1 {
		a, b := 0, 0
		if i < len(vec1) {
			a, _ = strconv.Atoi(vec1[i])
		}
		if j < len(vec2) {
			b, _ = strconv.Atoi(vec2[j])
		}
		if a > b {
			return 1
		}
		if b > a {
			return -1
		}
	}
	return 0
}

// @lc code=end
