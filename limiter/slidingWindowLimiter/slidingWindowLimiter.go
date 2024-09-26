package slidingWindowLimiter

import (
	"sync"
	"time"
)

// SlidingWindowLimiter 结构体实现滑动窗口限流算法
type SlidingWindowLimiter struct {
	mux            sync.Mutex
	counters       []uint64
	limit          uint64
	windowStart    time.Time
	windowDuration time.Duration
	interval       time.Duration
}

// NewSlidingWindowLimiter 构造函数初始化 SlidingWindowLimiter 实例。
func NewSlidingWindowLimiter(limit uint64, windowDuration, interval time.Duration) *SlidingWindowLimiter {
	buckets := int(windowDuration / interval)
	return &SlidingWindowLimiter{
		counters:       make([]uint64, buckets),
		limit:          limit,
		windowStart:    time.Now(),
		windowDuration: windowDuration,
		interval:       interval,
	}
}

// Allow 方法用于判断当前请求是否被允许，并实现滑动窗口的逻辑。
func (s *SlidingWindowLimiter) Allow() bool {
	s.mux.Lock()
	defer s.mux.Unlock()

	// 检查是否需要滑动窗口
	if time.Since(s.windowStart) > s.windowDuration {
		s.slidWindow()
	}

	now := time.Now()
	index := int((now.UnixNano()-s.windowStart.UnixNano())/s.interval.Nanoseconds()) % len(s.counters)

	if s.counters[index] < s.limit {
		s.counters[index]++
		return true
	}
	return false
}

// slideWindow 方法实现滑动窗口逻辑，移除最旧的时间段并重置计数器。
func (s *SlidingWindowLimiter) slidWindow() {
	// 滑动窗口，忽略最旧的时间段
	copy(s.counters, s.counters[1:])
	// 重置最后一个时间段的计数器
	s.counters[len(s.counters)-1] = 0
	// 更新窗口开始时间
	s.windowStart = time.Now()
}

/*
实现原理： 滑动窗口算法通过将时间分为多个小的时间段，每个时间段内维护一个独立的计数器。当一个请求到达时，它会被分配到当前时间所在的小时间段，并检查该时间段的计数器是否已达到限制。
如果未达到，则允许请求并增加计数；如果已达到，则拒绝请求。随着时间的推移，旧的时间段会淡出窗口，新的时间段会加入。

优点：
	相比固定窗口算法，滑动窗口算法能够更平滑地处理请求，避免瞬时高峰。
	可以提供更细致的流量控制。

缺点：
	实现相对复杂，需要维护多个计数器和时间索引。
	对内存和计算的要求更高。

滑动窗口算法适用于需要平滑流量控制的场景，尤其是在面对突发流量时，能够提供比固定窗口计数器更优的流量控制效果。


*/
