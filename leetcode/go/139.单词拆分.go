/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */
package main

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool, 0)
	for _, str := range wordDict {
		wordSet[str] = true
	}
	dp := make(map[int]bool, len(s)+1) //dp[i]表示s前i个字符能否拆分
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			//dp[i] = dp[j]&&check(s[j..i−1])
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// @lc code=end
