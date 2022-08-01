/*
 * @lc app=leetcode.cn id=316 lang=golang
 *
 * [316] 去除重复字母
 */

package main

// @lc code=start
func removeDuplicateLetters(s string) string {
	countMap := [256]int{}
	visited := [256]bool{}
	st := []byte{}
	for i := range s {
		c := int(s[i])
		countMap[c]++ // countMap[int(c)]=i
	}

	for i := range s {
		c := int(s[i])
		countMap[c]-- //每遍历过一个字符，都将对应的计数减一
		if visited[c] {
			continue
		}
		for len(st) > 0 {
			last := len(st) - 1
			lastCh := int(st[last])
			// 插入之前，和之前的元素比较一下大小
			// 如果字典序比前面的大，停止pop 前面的元素
			// 若之后不存在栈顶元素了，则停止 pop
			if lastCh <= c || countMap[lastCh] == 0 {
				break
			}
			// c < stack.top(); 弹出栈顶元素，并把该元素标记为不在栈中
			st = st[:last]
			visited[lastCh] = false
		}
		visited[c] = true
		st = append(st, byte(c))
	}
	return string(st)
}

// @lc code=end
