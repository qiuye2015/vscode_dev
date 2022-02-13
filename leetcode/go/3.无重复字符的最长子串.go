/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
package main

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	ret := 0
	l := 0 // 指向该无重复子串左边的起始位置
	count := make(map[byte]int, 0)
	for r := 0; r < len(s); r++ {
		// if left, ok := count[s[r]]; ok {
		// 	for ; l <= left; l++ { // l指向重复字符下一个字符
		// 		delete(count, s[l]) // 删除到重复字符之间的字符
		// 	}
		// }
		if left, ok := count[s[r]]; ok && left >= l {
			l = left + 1
		}
		count[s[r]] = r
		ret = max(ret, r-l+1)
	}
	return ret
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
