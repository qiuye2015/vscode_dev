/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
 */

package main

// @lc code=start
func maxCoins(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
	}
	// 计算开区间(i,j)得到得最大金币数
	range_best := func(i, j int) {
		res := 0
		// k是(i,j)区间内最后一个被戳的气球
		for k := i + 1; k < j; k++ {
			tmp := dp[i][k] + dp[k][j] + nums[i]*nums[k]*nums[j]
			if tmp > res {
				res = tmp
				// fmt.Printf("dp[%d][%d]=%d,k=%d\n", i, j, res, k)
			}
		}
		dp[i][j] = res
	}
	// 对每一个区间长度进行循环,长度从2开始
	for n := 2; n < len(nums); n++ {
		for i := 0; i < len(nums)-n; i++ {
			range_best(i, i+n)
		}
	}
	// for i := 0; i < len(nums)-2; i++ {
	// 	for j := i + 2; j < len(nums); j++ {
	// 		// 计算开区间(i,j)得到得最大金币数
	// 		// k是(i,j)区间内最后一个被戳的气球
	// 		for k := i + 1; k < j; k++ {
	// 			total := dp[i][k] + dp[k][j] + nums[i]*nums[k]*nums[j]
	// 			if total > dp[i][j] {
	// 				dp[i][j] = total
	// 				fmt.Printf("dp[%d][%d]=%d,k=%d\n", i, j, dp[i][j], k)
	// 			}
	// 		}
	// 	}
	// }
	return dp[0][len(nums)-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
