/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

package main

// @lc code=start
func isHappy(n int) bool {
	fast, slow := n, n
	first := true
	for first || fast != slow {
		first = false
		slow = bitSquareSum(slow)
		fast = bitSquareSum(fast)
		fast = bitSquareSum(fast)
	}
	return fast == 1
}
func bitSquareSum(n int) int {
	sum := 0
	for n > 0 {
		tmp := n % 10
		sum += tmp * tmp
		n = n / 10
	}
	return sum
}

// @lc code=end
