/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	ret := make([]int, 0)
	if len(s) == 0 || len(p) == 0 || len(s) < len(p) {
		return ret
	}
	record := make(map[byte]int, 0)
	for i := 0; i < len(p); i++ {
		record[p[i]]++
	}
	l, r := 0, 0
	for r < len(s) {
		ch := s[r]
		// 将右边界字符纳入窗口，更新其在 record 中的计数
		record[ch]--
		// 如果 record 中 ch 的计数小于 0，说明当前窗口中 ch 出现次数超过 p 中的次数，需要移动左边界
		for record[ch] < 0 {
			// 将左边界字符移出窗口，更新其在 record 中的计数
			record[s[l]]++
			// 左边界右移
			l++
		}
		// 当窗口大小等于 p 的长度时，找到一个符合条件的子串，记录起始索引 l
		if r-l+1 == len(p) {
			ret = append(ret, l)
		}
		// 右边界右移，扩大窗口
		r++
	}
	return ret
}

// @lc code=end

