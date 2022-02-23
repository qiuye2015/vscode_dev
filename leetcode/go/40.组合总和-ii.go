/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */
package main

import "sort"

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	ret := [][]int{}
	sort.Ints(candidates) // 排序是剪枝的前提

	var dfs func(cur []int, sum, start int)
	dfs = func(cur []int, sum, start int) {
		if sum == target {
			tmp := append([]int{}, cur...)
			ret = append(ret, tmp)
			return
		}
		for i := start; i < len(candidates); i++ {
			if sum+candidates[i] > target { // 剪枝，前提是候选数组已经有序，
				return
			}
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			cur = append(cur, candidates[i])
			dfs(cur, sum+candidates[i], i+1) // 由于每一个元素可以重复使用，下一轮搜索的起点依然是 i
			cur = cur[:len(cur)-1]
		}
	}
	dfs(make([]int, 0), 0, 0)
	return ret
}

// @lc code=end
