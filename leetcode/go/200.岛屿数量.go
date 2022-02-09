/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */
package main

// @lc code=start
func numIslands(grid [][]byte) int {
	ret := 0
	if len(grid) == 0 || len(grid[0]) == 0 {
		return ret
	}
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				checkIsland(i, j, &grid)
				ret++
			}
		}
	}
	return ret
}
func checkIsland(i, j int, grid *[][]byte) {
	if i < 0 || i >= len(*grid) || j < 0 || j >= len((*grid)[0]) {
		return
	}
	if (*grid)[i][j] != '1' {
		return
	}
	(*grid)[i][j] = '0' //表示已经访问过 <i,j>
	checkIsland(i+1, j, grid)
	checkIsland(i-1, j, grid)
	checkIsland(i, j+1, grid)
	checkIsland(i, j-1, grid)
}

// @lc code=end
