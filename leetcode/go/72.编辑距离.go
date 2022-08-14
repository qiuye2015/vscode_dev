/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

package main

// @lc code=start
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1) //DP[i,j]:前i个字符到前j个字符的最小操作数
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i //删除i个字符即可
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j //添加j个字符即可
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				w1Add := dp[i][j-1]
				w1Del := dp[i-1][j]
				w1Up := dp[i-1][j-1]
				dp[i][j] = 1 + min(w1Add, min(w1Del, w1Up))
			}
		}
	}
	return dp[m][n]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
