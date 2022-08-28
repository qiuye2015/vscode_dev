/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N皇后 II
 */

package main

// @lc code=start
func totalNQueens(n int) int {
	cols := make([]bool, n)    //列上是否有皇后
	diag1 := make([]bool, 2*n) //左上到右下是否有皇后 \  行下标与列下标之差相等
	diag2 := make([]bool, 2*n) //右上到左下是否有皇后 /  行下标与列下标之和相等
	res := 0
	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			res++
			return
		}
		for col := range cols {
			idx1, idx2 := row-col+n-1, row+col
			if cols[col] || diag1[idx1] || diag2[idx2] {
				continue
			}
			cols[col], diag1[idx1], diag2[idx2] = true, true, true
			dfs(row + 1)
			cols[col], diag1[idx1], diag2[idx2] = false, false, false
		}
	}
	dfs(0)
	return res
}

// @lc code=end
