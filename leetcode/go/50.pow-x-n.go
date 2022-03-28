/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */
package main

// @lc code=start
func myPow(x float64, n int) float64 {
	minus := n < 0
	if minus {
		n = -n
	}
	ret := float64(1)
	for n > 0 {
		if n&0x01 == 1 {
			ret *= x
		}
		x *= x
		n /= 2
	}
	if minus {
		ret = 1 / ret
	}
	return ret
}

// @lc code=end
