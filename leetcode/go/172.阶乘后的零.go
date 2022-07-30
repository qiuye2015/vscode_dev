/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 */

package main

// @lc code=start
func trailingZeroes(n int) int {
	count := 0
	for n > 0 {
		count += n / 5
		n /= 5
	}
	return count
}

// @lc code=end
