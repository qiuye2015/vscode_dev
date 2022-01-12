/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
package main

func twoSum(nums []int, target int) []int {
	record := make(map[int]int, 0)
	for idx, num := range nums {
		if val, ok := record[target-num]; ok {
			return []int{val, idx}
		}
		record[num] = idx
	}
	return nil
}

// @lc code=end
