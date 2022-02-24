/*
 * @lc app=leetcode.cn id=377 lang=golang
 *
 * [377] 组合总和 Ⅳ
 */
package main

// @lc code=start
func combinationSum4(nums []int, target int) int {
	record := map[int]int{}
	var dfs func(wanted int) int
	dfs = func(wanted int) int {
		if wanted < 0 {
			return 0
		} else if wanted == 0 {
			return 1
		}
		if cnt, exit := record[wanted]; exit {
			return cnt
		}

		cnt := 0
		for _, num := range nums {
			cnt += dfs(wanted - num)
		}
		record[wanted] = cnt
		return cnt
	}
	return dfs(target)
}

// @lc code=end
