/*
 * @lc app=leetcode.cn id=648 lang=golang
 *
 * [648] 单词替换
 */

package main

import (
	"strings"
)

// @lc code=start
func replaceWords(dictionary []string, sentence string) string {
	dictSet := map[string]bool{}
	for _, val := range dictionary {
		dictSet[val] = true
	}
	brokenString := strings.Fields(sentence)
	for i, str := range brokenString {
		for j := 1; j <= len(str); j++ {
			if dictSet[str[:j]] {
				brokenString[i] = str[:j]
				break
			}
		}
	}
	return strings.Join(brokenString, " ")
}

// @lc code=end
