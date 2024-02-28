/*
 * @lc app=leetcode.cn id=841 lang=golang
 *
 * 841.钥匙和房间
 */

package main

// @lc code=start
func canVisitAllRooms(rooms [][]int) bool {
	num := len(rooms)
	vis := make([]bool, num)
	var dfs func(x int)
	dfs = func(x int) {
		num--
		vis[x] = true
		for _, v := range rooms[x] {
			if !vis[v] {
				dfs(v)
			}
		}
	}
	dfs(0)
	return num == 0
}

// @lc code=end
