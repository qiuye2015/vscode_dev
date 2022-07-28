/*
 * @lc app=leetcode.cn id=1206 lang=golang
 *
 * [1206] 设计跳表
 */

package main

import (
	"math"
	"math/rand"
	"time"
)

// @lc code=start
type Skiplist struct {
	maxLevel int
	head     *Node
}

type Node struct {
	val   int
	count int
	next  []*Node
}

func Constructor() Skiplist {
	level := 16
	rand.Seed(time.Now().UnixNano())
	tail := &Node{
		val:   math.MaxInt,
		count: 0,
		next:  nil,
	}
	next := make([]*Node, level)
	for i := range next {
		next[i] = tail
	}
	head := &Node{
		val:   -1,
		count: 0,
		next:  next,
	}
	return Skiplist{
		maxLevel: level,
		head:     head,
	}
}

func (this *Skiplist) Search(target int) bool {
	curNode := this.head
	curLevel := 0
	for curLevel < this.maxLevel {
		if curNode.next[curLevel].val == target { // found, return true
			// fmt.Println(fmt.Sprintf("search %v, found", target))
			return true
		} else if curNode.next[curLevel].val < target {
			curNode = curNode.next[curLevel] // try next node
			continue
		}
		curLevel++ // try next level
	}
	// fmt.Println(fmt.Sprintf("search %v, not found", target))
	return false // not found
}

func (this *Skiplist) Add(num int) {
	// random pick a level
	addLevel := rand.Intn(this.maxLevel)
	curNode := this.head
	curLevel := 0
	needAdd := make([]*Node, 0)
	for curLevel < this.maxLevel {
		if curNode.next[curLevel].val == num { // already exist
			// fmt.Println(fmt.Sprintf("add %v already exist, count++", num))
			curNode.next[curLevel].count++
			return
		} else if curNode.next[curLevel].val < num {
			curNode = curNode.next[curLevel] // try next node
			continue
		}
		if curLevel >= addLevel {
			needAdd = append(needAdd, curNode)
		}
		curLevel++ // try next level
	}
	// num not exist, insert into SkipList
	newNode := &Node{
		val:   num,
		count: 1,
		next:  make([]*Node, this.maxLevel),
	}
	for l, v := range needAdd {
		newNode.next[l+addLevel] = v.next[l+addLevel]
		v.next[l+addLevel] = newNode
	}
	// fmt.Println(fmt.Sprintf("add %v successfully", num))
}

func (this *Skiplist) Erase(num int) bool {
	curNode := this.head
	curLevel := 0
	erased := false
	for curLevel < this.maxLevel {
		if curNode.next[curLevel].val == num {
			if curNode.next[curLevel].count > 1 {
				curNode.next[curLevel].count--
				return true
			}
			curNode.next[curLevel] = curNode.next[curLevel].next[curLevel]
			erased = true
			curLevel++
			continue
		} else if curNode.next[curLevel].val < num {
			curNode = curNode.next[curLevel]
			continue
		}
		curLevel++
	}
	// fmt.Println(fmt.Sprintf("erase %v result %v", num, erased))
	return erased
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
// @lc code=end
