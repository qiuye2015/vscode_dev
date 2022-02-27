/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */
package main

// @lc code=start
func permute(nums []int) [][]int {
	ret := [][]int{}
	var dfs func(cur, left []int)
	dfs = func(cur, left []int) {
		if len(left) == 0 {
			ret = append(ret, cur)
			return
		}
		for idx, num := range left {
			newLeft := append([]int{}, left[:idx]...)
			newLeft = append(newLeft, left[idx+1:]...)
			tmpCur := append(cur, num)
			dfs(tmpCur, newLeft)
		}
	}
	dfs(make([]int, 0), nums)
	return ret
}

// @lc code=end
