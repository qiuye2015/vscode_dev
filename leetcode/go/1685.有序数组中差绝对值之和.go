/*
 * @lc app=leetcode.cn id=1685 lang=golang
 *
 * [1685] 有序数组中差绝对值之和
 */

package main

import (
	"fmt"
)

// @lc code=start
func getSumAbsoluteDifferences(nums []int) []int {
	total, n := 0, len(nums)
	prefix := make([]int, n) //前缀和
	for i := 0; i < n; i++ {
		total += nums[i]
		prefix[i] = total
	}
	fmt.Println(prefix)
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		before := nums[i]*(i+1) - prefix[i]
		after := (total - prefix[i]) - (nums[i] * (n - 1 - i))
		ret[i] = before + after
	}
	return ret
}

// @lc code=end
