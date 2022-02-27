/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */
package main

import "sort"

// @lc code=start
func permuteUnique(nums []int) [][]int {
	ret := [][]int{}
	visited := make([]bool, len(nums))
	sort.Ints(nums)
	var dfs func(cur []int)
	dfs = func(cur []int) {
		if len(cur) == len(nums) {
			tmp := append([]int{}, cur...)
			ret = append(ret, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}
			// 对于重复数的集合，一定是从左往右逐个填入
			// 前面的未填入的话，当前的也不能填入
			if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
				continue
			}
			visited[i] = true
			cur = append(cur, nums[i])
			dfs(cur)
			cur = cur[:len(cur)-1]
			visited[i] = false
		}
	}
	dfs(make([]int, 0))
	return ret
}

// @lc code=end
