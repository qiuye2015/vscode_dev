/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */
package main

import (
	"sort"
)

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	ret := [][]int{{}}
	cur := []int{}
	sort.Ints(nums)

	var dfs func(nums, cur []int)
	dfs = func(nums, cur []int) {
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			cur = append(cur, nums[i])
			tmp := append([]int{}, cur...)
			ret = append(ret, tmp)

			dfs(nums[i+1:], cur)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(nums, cur)
	return ret
}

// @lc code=end
