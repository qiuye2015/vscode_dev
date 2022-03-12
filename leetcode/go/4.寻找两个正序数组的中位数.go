/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */
package main

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	if n%2 == 0 {
		left := findKth(nums1, nums2, n/2)
		right := findKth(nums1, nums2, n/2+1)
		return (left + right) / 2.0
	} else {
		return findKth(nums1, nums2, n/2+1)
	}
}

// 寻找第K大的数（k从1开始）
func findKth(nums1, nums2 []int, k int) float64 {
	if len(nums1) > len(nums2) { // 确保nums1长度较小
		return findKth(nums2, nums1, k)
	}
	if len(nums1) == 0 {
		return float64(nums2[k-1])
	}
	if k == 1 {
		return float64(minInt2(nums1[0], nums2[0]))
	}
	idx1 := minInt2(k/2, len(nums1))
	idx2 := k - k/2

	if nums1[idx1-1] <= nums2[idx2-1] {
		return findKth(nums1[idx1:], nums2, k-idx1)
	} else {
		return findKth(nums1, nums2[idx2:], k-idx2)
	}
}
func minInt2(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// @lc code=end
