/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

package main

// @lc code=start
func climbStairs(n int) int {
	p, q := 1, 1
	for i := 2; i <= n; i++ {
		t := p + q
		p = q
		q = t
	}
	return q
}

// @lc code=end
