/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

package main

// @lc code=start
// func partition(s string) [][]string {
// 	res := [][]string{}
// 	var dfs func(cur []string, start int)
// 	dfs = func(cur []string, start int) {
// 		if start == len(s) {
// 			tmp := append([]string{}, cur...)
// 			res = append(res, tmp)
// 			return
// 		}
// 		for i := start; i < len(s); i++ {
// 			tmpStr := s[start : i+1]
// 			if isPalindrome(tmpStr) {
// 				cur = append(cur, tmpStr)
// 				dfs(cur, i+1)
// 				cur = cur[:len(cur)-1]
// 			}
// 		}
// 	}
// 	dfs([]string{}, 0)
// 	return res
// }

// func isPalindrome(s string) bool {
// 	i, j := 0, len(s)-1
// 	for i < j {
// 		if s[i] != s[j] {
// 			return false
// 		}
// 		i, j = i+1, j-1
// 	}
// 	return true
// }
func partition(s string) [][]string {
	n := len(s)
	res := [][]string{}
	dp := make([][]bool, n) //dp[i][j]表示s[i][j]是否是回文
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for j := 0; j < n; j++ {
		for i := 0; i <= j; i++ {
			// j=i   一个字符 a
			// j-i=1 两个字符 aa
			// j-i=2 三个字符 aba
			dp[i][j] = s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1])
		}
	}
	var dfs func(cur []string, start int)
	dfs = func(cur []string, start int) {
		if start == n {
			tmp := append([]string{}, cur...)
			res = append(res, tmp)
			return
		}
		for i := start; i < n; i++ {
			if dp[start][i] {
				cur = append(cur, s[start:i+1])
				dfs(cur, i+1)
				cur = cur[:len(cur)-1]
			}
		}
	}
	dfs([]string{}, 0)
	return res
}

// @lc code=end
