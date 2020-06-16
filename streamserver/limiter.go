// 视频流控

package main

import (
	"log"
)

// ConnLimiter 连接限制器
type ConnLimiter struct {
	concurrentConn int //连接的个数
	bucket         chan int
}

// NewConnLimiter 初始化连接限制器
func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}

// GetConn 获取连接
func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reached the rate limitation.")
		return false
	}

	cl.bucket <- 1
	return true
}

// ReleaseConn 释放连接
func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket
	log.Printf("New connction coming: %d", c)
}
