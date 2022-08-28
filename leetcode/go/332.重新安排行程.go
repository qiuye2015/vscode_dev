/*
 * @lc app=leetcode.cn id=332 lang=golang
 *
 * [332] 重新安排行程
 */

package main

import "sort"

// @lc code=start
func findItinerary(tickets [][]string) []string {
	//当我们遍历完一个节点所连的所有节点后，我们才将该节点入栈（即逆序入栈）
	//整个图最多存在一个死胡同(出度和入度相差1），
	//且这个死胡同一定是最后一个访问到的，否则无法完成一笔画
	m := map[string][]string{}
	res := []string{}
	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		m[src] = append(m[src], dst)
	}
	for src := range m {
		sort.Strings(m[src])
	}
	var dfs func(cur string)
	dfs = func(cur string) {
		for {
			if v, ok := m[cur]; !ok || len(v) == 0 {
				break
			}
			tmp := m[cur][0]
			m[cur] = m[cur][1:]
			dfs(tmp)
		}
		res = append(res, cur)
	}
	dfs("JFK")
	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}

// @lc code=end
