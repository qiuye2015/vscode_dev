/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */
package main

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	if len(matrix) == 0 {
		return res
	}
	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for {
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		up++
		if up > down {
			break
		}
		for j := up; j <= down; j++ {
			res = append(res, matrix[j][right])
		}
		right--
		if right < left {
			break
		}
		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}
		down--
		if down < up {
			break
		}
		for j := down; j >= up; j-- {
			res = append(res, matrix[j][left])
		}
		left++
		if left > right {
			break
		}
	}
	return res
}

// @lc code=end
