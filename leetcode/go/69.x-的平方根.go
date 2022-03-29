/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */
package main

// @lc code=start
func mySqrt(x int) int {
	l, r := 0, x
	ret := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			l = mid + 1
			ret = mid
		} else {
			r = mid - 1
		}
	}
	return ret
}

// @lc code=end
