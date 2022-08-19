/*
 * @lc app=leetcode.cn id=1901 lang=golang
 *
 * [1901] 找出顶峰元素 II
 */

package main

// @lc code=start
func findPeakGrid(mat [][]int) []int {
	low, high := 0, len(mat)-1
	for low < high {
		mid := low + (high-low)>>1
		idx := getMaxIndex(mat[mid])
		if mat[mid][idx] > mat[mid+1][idx] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	idx := getMaxIndex(mat[low])
	return []int{low, idx}
}

//获取单行最大值位置
func getMaxIndex(nums []int) int {
	max, idx := 0, 0
	for i, num := range nums {
		if num > max {
			max = num
			idx = i
		}
	}
	return idx
}

// @lc code=end
