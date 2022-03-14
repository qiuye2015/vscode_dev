/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */
package main

import (
	"container/list"
)

// @lc code=start

type Item struct {
	Key int
	Val int
}
type LRUCache struct {
	Capacity  int
	ItemsList *list.List
	ItemsMap  map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity:  capacity,
		ItemsList: list.New(),
		ItemsMap:  map[int]*list.Element{},
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.ItemsMap[key]; ok {
		this.ItemsList.MoveToFront(node)
		return node.Value.(Item).Val
	}
	// value := getValueFromDB(key)
	// l.Put(key, value)
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	//lock&unlock
	if node, ok := this.ItemsMap[key]; ok {
		node.Value = Item{
			Key: key,
			Val: value,
		}
		this.ItemsList.MoveToFront(node)
		return
	}
	if len(this.ItemsMap) == this.Capacity {
		back := this.ItemsList.Back()
		delete(this.ItemsMap, back.Value.(Item).Key)
		this.ItemsList.Remove(back)
	}
	node := this.ItemsList.PushFront(Item{
		Key: key,
		Val: value,
	})
	this.ItemsMap[key] = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
