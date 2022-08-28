/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

package main

// @lc code=start
func solveNQueens(n int) [][]string {
	res := [][]string{}
	cols := make([]bool, n)
	diag1 := make([]bool, 2*n)
	diag2 := make([]bool, 2*n)
	rows := make([]int, 0) //写入结果
	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			tmp := genBoard(rows)
			res = append(res, tmp)
			return
		}
		for col := range cols {
			idx1, idx2 := row-col+n-1, row+col
			if cols[col] || diag1[idx1] || diag2[idx2] {
				continue
			}
			cols[col], diag1[idx1], diag2[idx2] = true, true, true
			rows = append(rows, col)
			dfs(row + 1)
			cols[col], diag1[idx1], diag2[idx2] = false, false, false
			rows = rows[:len(rows)-1]
		}
	}
	dfs(0)
	return res
}
func genBoard(rows []int) []string {
	n := len(rows)
	ans := []string{}
	for i := 0; i < n; i++ {
		line := ""
		for j := 0; j < n; j++ {
			if rows[i] == j {
				line += "Q"
			} else {
				line += "."
			}
		}
		ans = append(ans, line)
	}
	return ans
}

// @lc code=end
