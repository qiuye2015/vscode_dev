/*
 * @lc app=leetcode.cn id=670 lang=golang
 *
 * [670] 最大交换
 */

package main

import (
	"strconv"
)

// @lc code=start
func maximumSwap(num int) int {
	s := []byte(strconv.Itoa(num))
	max, maxIdx := -1, -1
	i, j := 0, 0
	for k := len(s) - 1; k >= 0; k-- {
		digit := int(s[k] - '0')
		if digit > max {
			max = digit
			maxIdx = k
		}
		if digit < max {
			i, j = maxIdx, k
		}
	}
	s[i], s[j] = s[j], s[i]
	ret, _ := strconv.Atoi(string(s))
	return ret
}

// @lc code=end
