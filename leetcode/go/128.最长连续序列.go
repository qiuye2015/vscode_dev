/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

package main

// @lc code=start
func longestConsecutive(nums []int) int {
	m := make(map[int]bool, 0)
	for _, num := range nums {
		m[num] = true
	}
	ret := 0
	for num := range m {
		if m[num-1] { //避免重复计算
			continue
		}
		cur := num
		for m[cur+1] {
			cur++
		}
		cnt := cur - num + 1
		if cnt > ret {
			ret = cnt
		}
	}
	return ret
}

// @lc code=end
