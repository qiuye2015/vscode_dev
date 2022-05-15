/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
package main

// @lc code=start
func longestPalindrome(s string) string {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	start := 0
	maxLen := 1
	for i := 0; i < len(s)-1; i++ { // i为中心位置
		oddLen := expandCentor(s, i, i)
		evenLen := expandCentor(s, i, i+1)
		curMaxLen := max(oddLen, evenLen)
		if curMaxLen > maxLen {
			maxLen = curMaxLen
			start = i - (maxLen-1)/2
		}
	}
	return s[start : start+maxLen]
}

// 以left，right为中心向两边扩展，能取到最长的回文子串
func expandCentor(s string, left, right int) int {
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			break
		}
		left--
		right++
	}
	// 此时 s[left]!=s[right]
	// 回文子串[left+1,right-1] ==>
	// right-1 - (left +1) + 1
	return right - left - 1
}

// @lc code=end
