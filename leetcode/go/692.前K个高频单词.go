/*
 * @lc app=leetcode.cn id=692 lang=golang
 *
 * [692] 前K个高频单词
 */

package main

import (
	"sort"
)

// @lc code=start
func topKFrequent(words []string, k int) []string {
	cnt := map[string]int{}
	for _, w := range words {
		cnt[w]++
	}
	uniqueWords := []string{}
	for k := range cnt {
		uniqueWords = append(uniqueWords, k)
	}
	sort.Slice(uniqueWords, func(i, j int) bool {
		s, t := uniqueWords[i], uniqueWords[j]
		if cnt[s] == cnt[t] {
			return s < t
		}
		return cnt[s] > cnt[t]
	})
	return uniqueWords[:k]
}

// @lc code=end
