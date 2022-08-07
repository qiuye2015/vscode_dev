/*
 * @lc app=leetcode.cn id=460 lang=golang
 *
 * [460] LFU 缓存
 */

package main

import (
	"container/list"
)

// @lc code=start
type Item struct {
	key, val, freq int
}
type LFUCache struct {
	cap, minFreq int
	list         map[int]*list.List    // freq --> Items
	cache        map[int]*list.Element // key  --> index
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:   capacity,
		list:  make(map[int]*list.List),
		cache: make(map[int]*list.Element),
	}
}

func (this *LFUCache) Get(key int) int {
	if e, ok := this.cache[key]; ok {
		e = this.delAndAdd(e)
		return e.Value.(*Item).val
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}
	if e, ok := this.cache[key]; ok {
		e = this.delAndAdd(e)
		e.Value.(*Item).val = value
	} else {
		if len(this.cache) == this.cap {
			e := this.list[this.minFreq].Back()
			this.list[this.minFreq].Remove(e)
			delete(this.cache, e.Value.(*Item).key)
		}
		this.minFreq = 1
		newItem := &Item{
			key:  key,
			val:  value,
			freq: 1,
		}
		if _, ok := this.list[this.minFreq]; !ok {
			this.list[this.minFreq] = list.New()
		}
		e := this.list[this.minFreq].PushFront(newItem)
		this.cache[key] = e
	}

}
func (this *LFUCache) delAndAdd(e *list.Element) *list.Element {
	freq := e.Value.(*Item).freq
	key := e.Value.(*Item).key
	//delete
	this.list[freq].Remove(e)
	if this.minFreq == freq && this.list[freq].Len() == 0 {
		this.minFreq++
	}
	newFreq := freq + 1
	e.Value.(*Item).freq = newFreq
	if _, ok := this.list[newFreq]; !ok {
		this.list[newFreq] = list.New()
	}
	//add
	e = this.list[newFreq].PushFront(e.Value)
	this.cache[key] = e
	return e
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
