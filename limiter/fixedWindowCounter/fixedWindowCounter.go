package fixedWindowCounter

import (
	"sync"
	"time"
)

// FixedWindowCounter 结构体实现固定窗口计数器限流算法。
// mu 用于同步访问，保证并发安全。
// count 记录当前时间窗口内的请求数量。
// limit 是时间窗口内允许的最大请求数量。
// window 记录当前时间窗口的开始时间。
// duration 是时间窗口的持续时间。
type FixedWindowCounter struct {
	mux      sync.Mutex
	count    int
	limit    int
	window   time.Time
	duration time.Duration
}

// NewFixedWindowCounter 构造函数初始化 FixedWindowCounter 实例。
func NewFixedWindowCounter(limit int, duration time.Duration) *FixedWindowCounter {
	return &FixedWindowCounter{
		limit:    limit,
		window:   time.Now(), // 设置当前时间作为窗口的开始时间。
		duration: duration,   // 设置时间窗口的持续时间。
	}
}

// Allow 方法用于判断当前请求是否被允许。
// 首先通过互斥锁保证方法的原子性。
func (f *FixedWindowCounter) Allow() bool {
	f.mux.Lock()
	defer f.mux.Unlock()
	now := time.Now()

	// 如果当前时间超过了窗口的结束时间，重置计数器和窗口开始时间。
	if now.After(f.window.Add(f.duration)) {
		f.count = 0
		f.window = now
	}

	// 如果当前计数小于限制，则增加计数并允许请求。
	if f.count < f.limit {
		f.count++
		return true
	}
	// 如果计数达到限制，则拒绝请求。
	return false
}

/*
实现原理： 固定窗口计数器算法通过设置一个固定的时间窗口（例如每分钟）和一个在这个窗口内允许的请求数量限制（例如10个请求）。在每个时间窗口开始时，计数器重置为零，随着请求的到来，计数器递增。当计数器达到限制时，后续的请求将被拒绝，直到窗口重置。

优点：
	实现简单直观。
	容易理解和实现。
	可以保证在任何给定的固定时间窗口内，请求的数量不会超过设定的阈值。

缺点：
	在窗口切换的瞬间可能会有请求高峰，因为计数器重置可能导致大量请求几乎同时被处理。
	无法平滑地处理突发流量，可能导致服务体验不佳。

固定窗口计数器算法适用于请求分布相对均匀的场景，但在请求可能在短时间内集中到达的场景下，可能需要考虑更复杂的限流算法，如滑动窗口或令牌桶算法。


*/
