/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

package main

// @lc code=start
func letterCombinations(digits string) []string {
	m := map[byte][]string{
		'0': []string{"0"},
		'1': []string{"1"},
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}
	ret := []string{}
	if len(digits) == 0 {
		return ret
	}
	var dfs func(cur string, index int)
	dfs = func(cur string, index int) {
		if index == len(digits) {
			ret = append(ret, cur)
			return
		}
		letters := m[digits[index]]
		for _, ch := range letters {
			dfs(cur+ch, index+1)
		}
	}
	dfs("", 0)
	return ret
}

// @lc code=end
