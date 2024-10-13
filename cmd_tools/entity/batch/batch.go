// Package batch 提供无阻塞多并发功能。
// tRPC-Go 框架提供的 trpc.GoAndWait，或 sync.WaitGroup 都是同时等待并发的函数，必然是以最长执行时间的函数为等待时间。
// 这样其他提前结束的函数势必会浪费大量的时间。
// 而为了提升并发效率，预期要一个函数执行完成后继续执行其他等待执行的函数（要执行的函数数量远远大于同时并发的数量）。
// 同时为保证最后一次并发的函数正常执行结束，只需要在结束时等待一次即可。
package batch

import (
	"time"
)

// Group 定义无阻塞多并发功能的数据结构。
type Group struct {
	chn chan bool
}

// New 创建一个无阻塞多并发功能的实例，cnt 为空闲 goroutine 的数量。
func New(cnt int) *Group {
	return &Group{
		chn: make(chan bool, cnt),
	}
}

// Add 添加一个并发任务，若有空闲则立即返回，若无空闲则等待。
// 注意：不要在协程中调用本方法等待添加一个并发任务，可能会造成协程数量暴增。
func (g *Group) Add() {
	g.chn <- true
}

// Done 标记结束一个并发任务
func (g *Group) Done() {
	<-g.chn
}

// Wait 等待所有 goroutine 都结束。只需要在等待所有并发任务结束时调用一次即可。
func (g *Group) Wait() {
	for len(g.chn) != 0 {
		time.Sleep(time.Microsecond)
	}
}
