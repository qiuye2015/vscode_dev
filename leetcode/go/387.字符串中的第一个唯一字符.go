/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */

package main

// @lc code=start
func firstUniqChar(s string) int {
	record := [26]int{}
	for _, c := range s {
		record[c-'a']++
	}
	for i, c := range s {
		if record[c-'a'] == 1 {
			return i
		}
	}
	return -1
}

// @lc code=end
