/*
 * @lc app=leetcode.cn id=491 lang=golang
 *
 * [491] 递增子序列
 */

package main

// @lc code=start
func findSubsequences(nums []int) [][]int {
	res := [][]int{}
	var dfs func(cur []int, start int)
	dfs = func(cur []int, start int) {
		if len(cur) > 1 {
			tmp := append([]int{}, cur...)
			res = append(res, tmp)
		}
		vis := map[int]bool{}
		for i := start; i < len(nums); i++ {
			if (start > 0 && nums[i] < nums[start-1]) || vis[nums[i]] {
				continue
			}
			vis[nums[i]] = true
			cur = append(cur, nums[i])
			dfs(cur, i+1)
			cur = cur[:len(cur)-1]
			// dfs(append(cur, nums[i]), i+1)
		}
	}
	dfs([]int{}, 0)
	return res
}

// @lc code=end
