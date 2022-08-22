/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为 K 的子数组
 */

package main

// @lc code=start
func subarraySum(nums []int, k int) int {
	prefix := map[int]int{0: 1} // 和为0的次数为1
	sum, count := 0, 0          // sum: nums0到i的和
	for _, num := range nums {
		sum += num
		count += prefix[sum-k]
		prefix[sum]++
	}
	return count
}

// @lc code=end
