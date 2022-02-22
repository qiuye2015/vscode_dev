/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */
package main

// @lc code=start
func combine(n int, k int) [][]int {
	ret := [][]int{}
	cur := []int{}
	var dfs func(start int, cur []int)
	dfs = func(start int, cur []int) {
		if len(cur)+(n-start+1) < k { // 剪枝：可有可无
			// cur长度加上区间还未访问的区间长度小于k，不可能构造出长度为 k 的 temp
			return
		}
		if len(cur) == k {
			// tmp := append([]int{}, cur...)
			tmp := make([]int, k)
			copy(tmp, cur)

			ret = append(ret, tmp)
			return
		}
		for i := start; i <= n; i++ {
			cur = append(cur, i) // 考虑选择当前位
			dfs(i+1, cur)
			cur = cur[:len(cur)-1] // 考虑不选择当前位置
		}
		// // 考虑选择当前位
		// cur = append(cur, start)
		// dfs(start+1, cur)
		// cur = cur[:len(cur)-1]
		// // 考虑不选择当前位置
		// dfs(start+1, cur)
	}
	dfs(1, cur)
	return ret
}

// @lc code=end
