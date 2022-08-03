/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

package main

// @lc code=start
func lengthOfLIS(nums []int) int {
	tail := []int{} //tail[i]表示:长度为i + 1的所有上升子序列的结尾的最小值
	idx := 0
	for _, num := range nums {
		idx = binarySearch(tail, num)
		if idx == len(tail) { //新数num，如果这个数严格大于有序数组tail的最后一个元素，就把num放在有序数组tail的后面
			tail = append(tail, num)
		} else { // 如果有序数组tail中存在大于(等于)num的元素，找到第1个，让它变小，
			tail[idx] = num // 这样我们就找到了一个结尾更小的相同长度的上升子序列
		}
	}
	return len(tail)
}
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1 //
	for left <= right {           //
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1 //
		}
	}
	return left //返回左边界
}

// @lc code=end
