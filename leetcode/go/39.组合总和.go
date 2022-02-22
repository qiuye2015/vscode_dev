/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */
package main

import "sort"

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	ret := [][]int{}
	sort.Ints(candidates)

	var dfs func([]int, int, int)
	dfs = func(cur []int, sum, start int) {
		if sum > target {
			return
		}
		if sum == target {
			tmp := append([]int{}, cur...)
			ret = append(ret, tmp)
			return
		}
		for i := start; i < len(candidates); i++ {
			if i > 0 && candidates[i] == candidates[i-1] {
				continue
			}
			cur = append(cur, candidates[i])
			dfs(cur, sum+candidates[i], i)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(make([]int, 0), 0, 0)
	return ret
}

// @lc code=end
