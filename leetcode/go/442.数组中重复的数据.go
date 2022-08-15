/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */

package main

// @lc code=start
// func findDuplicates(nums []int) []int {
// 	ret := []int{}
// 	for _, num := range nums {
// 		target := abs(num) - 1 //nums[i]需要放入nums[i]-1的位置（1-->0）
// 		if nums[target] < 0 {
// 			ret = append(ret, abs(num))
// 		} else {
// 			nums[target] *= -1
// 		}
// 	}
// 	return ret
// }
// func abs(a int) int {
// 	if a < 0 {
// 		return -a
// 	}
// 	return a
// }

func findDuplicates(nums []int) []int {
	ret := []int{}
	for i := range nums {
		//nums[i]放到nums[i]-1的位置
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := range nums {
		if nums[i] != i+1 {
			ret = append(ret, nums[i])
		}
	}
	return ret
}

// @lc code=end
