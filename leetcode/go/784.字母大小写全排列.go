/*
 * @lc app=leetcode.cn id=784 lang=golang
 *
 * [784] 字母大小写全排列
 */
package main

// @lc code=start

func letterCasePermutation(s string) []string {
	ret := []string{}
	var dfs func(cur []byte, start int)
	dfs = func(cur []byte, start int) {
		ret = append(ret, string(cur[:]))

		for i := start; i < len(s); i++ {
			if cur[i] >= 'a' && cur[i] <= 'z' {
				cur[i] = cur[i] - 32
				dfs(cur, i+1)
				cur[i] = cur[i] + 32
			} else if cur[i] >= 'A' && cur[i] <= 'Z' {
				cur[i] = cur[i] + 32
				dfs(cur, i+1)
				cur[i] = cur[i] - 32
			}
		}
	}
	dfs([]byte(s), 0)
	return ret
}

// @lc code=end
