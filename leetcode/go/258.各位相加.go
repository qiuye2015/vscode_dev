/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] 各位相加
 */

package main

// @lc code=start
func addDigits(num int) int {
	// if num < 10 {
	// 	return num
	// }
	// ret := num % 9
	// if ret == 0 {
	// 	return 9
	// } else {
	// 	return ret
	// }
	return (num-1)%9 + 1
}

// @lc code=end
