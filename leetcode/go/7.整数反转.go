/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

package main

import "math"

// @lc code=start
func reverse(x int) int {
	ret := 0
	for x != 0 {
		ret = ret*10 + x%10
		if ret > math.MaxInt32 || ret < math.MinInt32 {
			return 0
		}
		x = x / 10
	}
	return ret
}

// @lc code=end
