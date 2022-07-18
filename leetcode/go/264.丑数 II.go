/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

package main

// @lc code=start
func nthUglyNumber(n int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	if n == 1 {
		return 1
	}
	uglyVec := make([]int, n)
	uglyVec[0] = 1
	idx2, idx3, idx5 := 0, 0, 0
	for i := 1; i < n; i++ {
		new2 := uglyVec[idx2] * 2
		new3 := uglyVec[idx3] * 3
		new5 := uglyVec[idx5] * 5
		tmp := min(min(new2, new3), new5)
		uglyVec[i] = tmp
		if tmp == new2 {
			idx2++
		}
		if tmp == new3 {
			idx3++
		}
		if tmp == new5 {
			idx5++
		}
	}
	return uglyVec[n-1]
}

// @lc code=end
