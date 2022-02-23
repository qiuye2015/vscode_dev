/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 */
package main

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	ret := [][]int{}
	cur := []int{}
	var dfs func(int, int)
	dfs = func(start, sum int) {
		if sum == n && len(cur) == k {
			tmp := append([]int{}, cur...)
			ret = append(ret, tmp)
			return
		}

		for i := start; i <= 9; i++ {
			if sum+i > n || len(cur) > k {
				return
			}
			cur = append(cur, i)
			dfs(i+1, sum+i)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(1, 0)
	return ret
}

// @lc code=end
