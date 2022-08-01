/*
 * @lc app=leetcode.cn id=1081 lang=golang
 *
 * [1081] 不同字符的最小子序列
 */

package main

// @lc code=start
func smallestSubsequence(s string) string {
	visited := [256]bool{}
	countMap := [256]int{}
	stk := []byte{}
	for i := range s {
		c := int(s[i])
		countMap[c]++
	}
	for i := range s {
		c := int(s[i])
		countMap[c]--
		if visited[c] {
			continue
		}
		for len(stk) > 0 {
			last := len(stk) - 1
			lastCh := int(stk[last])
			if c >= lastCh || countMap[lastCh] == 0 {
				break
			}
			stk = stk[:last]
			visited[lastCh] = false
		}
		visited[c] = true
		stk = append(stk, byte(c))
	}
	return string(stk)
}

// @lc code=end
