/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */
package main

// @lc code=start
type Trie struct {
	IsEnd bool
	Son   [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	cur := this
	for _, ch := range word {
		idx := ch - 'a'
		if cur.Son[idx] == nil {
			cur.Son[idx] = &Trie{}
		}
		cur = cur.Son[idx]
	}
	cur.IsEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.searchPrefix(word)
	return node != nil && node.IsEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.searchPrefix(prefix) != nil
}

func (this *Trie) searchPrefix(prefix string) *Trie {
	cur := this
	for _, ch := range prefix {
		idx := ch - 'a'
		if cur.Son[idx] == nil {
			return nil
		}
		cur = cur.Son[idx]
	}
	return cur
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
