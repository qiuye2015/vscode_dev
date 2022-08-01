/*
 * @lc app=leetcode.cn id=321 lang=golang
 *
 * [321] 拼接最大数
 */

package main

// @lc code=start
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	n, m := len(nums1), len(nums2)
	fromFirst := min(n, k)
	fromSecond := k - fromFirst
	bestResult := []int{}
	for fromFirst >= 0 && fromSecond <= min(k, m) {
		a := maxNumSingle(nums1, fromFirst)
		b := maxNumSingle(nums2, fromSecond)
		merged := merge(a, b)
		if len(bestResult) == 0 || greater(merged, bestResult, 0, 0) {
			bestResult = merged
		}
		fromFirst--
		fromSecond++
	}
	return bestResult
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type intStack []int

func (s intStack) peek() int {
	return s[len(s)-1]
}
func (s *intStack) pop() int {
	last := len(*s) - 1
	ret := (*s)[last]
	*s = (*s)[:last]
	return ret
}
func (s *intStack) push(x int) {
	*s = append(*s, x)
}

func maxNumSingle(nums []int, k int) []int {
	stack := make(intStack, 0, k)
	n := len(nums)
	for i, num := range nums {
		left := n - i
		for len(stack) > 0 && num > stack.peek() &&
			left+len(stack) > k {
			stack.pop()
		}
		if len(stack) < k {
			stack.push(num)
		}
	}
	return stack
}

// a[i]>b[j]
func greater(a, b []int, i, j int) bool {
	for ; i < len(a) && j < len(b); i, j = i+1, j+1 {
		if a[i] != b[j] {
			return a[i] > b[j]
		}
	}
	return i != len(a)
}
func merge(a, b []int) []int {
	n, m := len(a), len(b)
	var i, j int
	res := make([]int, n+m)
	for k := 0; k < m+n; k++ {
		if greater(a, b, i, j) {
			res[k] = a[i]
			i++
		} else {
			res[k] = b[j]
			j++
		}
	}
	return res
}

// @lc code=end
