/*
 * @lc app=leetcode.cn id=223 lang=golang
 *
 * [223] 矩形面积
 */

package main

// @lc code=start
func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	area1 := (ax2 - ax1) * (ay2 - ay1)
	area2 := (bx2 - bx1) * (by2 - by1)
	sum := area1 + area2
	left, right := max(ax1, bx1), min(ax2, bx2)
	up, down := min(ay2, by2), max(ay1, by1)
	if right > left && up > down {
		sum -= (right - left) * (up - down) //overlap
	}
	return sum
}

// @lc code=end
