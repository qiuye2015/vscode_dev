/*
 * @lc app=leetcode.cn id=535 lang=golang
 *
 * [535] TinyURL 的加密与解密
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// @lc code=start
type Codec struct {
	s2l    map[string]string
	l2s    map[string]string
	prefix string
	str64  string
	strCnt int
}

func Constructor() Codec {
	rand.Seed(time.Now().UnixNano())
	return Codec{
		s2l:    make(map[string]string, 0),
		l2s:    make(map[string]string, 0),
		prefix: "http://tinyurl.com/",
		str64:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
		strCnt: 6,
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	for this.l2s[longUrl] == "" {
		cur := make([]byte, this.strCnt)
		for i := 0; i < this.strCnt; i++ {
			idx := rand.Intn(len(this.str64))
			cur[i] = this.str64[idx]
		}
		surl := fmt.Sprintf("%s%s", this.prefix, string(cur))
		if _, ok := this.s2l[surl]; ok {
			continue
		}
		this.l2s[longUrl] = surl
		this.s2l[surl] = longUrl
	}
	return this.l2s[longUrl]
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	if url, ok := this.s2l[shortUrl]; ok {
		return url
	}
	return ""
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
// @lc code=end
