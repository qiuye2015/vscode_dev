/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

package main

import (
	"math"
)

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	sum := 0
	ret := math.MaxInt
	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			cnt := right - left + 1
			if cnt < ret {
				ret = cnt
			}
			sum -= nums[left]
			left++
		}
		right++
	}
	if ret == math.MaxInt {
		return 0
	}
	return ret
}

// func minSubArrayLen(target int, nums []int) int {
// 	n := len(nums)
// 	if n == 0 {
// 		return 0
// 	}
// 	ret := math.MaxInt
// 	sums := make([]int, n+1) //前缀和: sums[i]表示从nums[0]到nums[i−1]的元素和
// 	// sums[0] = 0 意味着前 0 个元素的前缀和为 0
// 	// sums[1] = A[0] 前 1 个元素的前缀和为 A[0]
// 	for i := 1; i <= n; i++ {
// 		sums[i] = sums[i-1] + nums[i-1]
// 	}
// 	for i := 1; i <= n; i++ {
// 		s := sums[i-1] + target
// 		bound := sort.SearchInts(sums, s) // lower_bound:二分查找得到大于或等于s的最小下标bound
// 		cnt := bound - (i - 1)
// 		if bound <= n && cnt < ret {
// 			ret = cnt
// 		}
// 	}

// 	if ret == math.MaxInt {
// 		return 0
// 	}
// 	return ret
// }

// @lc code=end
