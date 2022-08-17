/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

package main

// @lc code=start
func minWindow(s string, t string) string {
	needCnt := len(t)
	counter := make(map[byte]int)
	for i := range t {
		counter[t[i]]++
	}
	if needCnt > len(s) {
		return ""
	}
	left, right := 0, 0
	ret := string(make([]byte, len(s)))
	for right < len(s) {
		if cnt, ok := counter[s[right]]; ok {
			if cnt > 0 { //这个是必须的，总数才--
				needCnt--
			}
			counter[s[right]]--
		}
		for needCnt <= 0 { //[left,right] 包含了所有的t
			sLen := right - left + 1
			if sLen <= len(ret) {
				ret = s[left : right+1]
			}
			if _, ok := counter[s[left]]; ok {
				counter[s[left]]++
				if counter[s[left]] > 0 {
					needCnt++
				}
			}
			left++
		}
		right++
	}
	if right == len(s) && left == 0 {
		return ""
	}
	return ret
}

// @lc code=end
