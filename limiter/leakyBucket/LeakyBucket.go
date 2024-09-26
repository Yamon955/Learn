package leakyBucket

import (
	"fmt"
	"time"
)

// LeakyBucket 结构体，包含请求队列
type LeakyBucket struct {
	queue chan struct{}
}

func NewLeakBucket(capacity int) *LeakyBucket {
	return &LeakyBucket{
		queue: make(chan struct{}, capacity),
	}
}

// Push 将请求放入队列，如果队列满了，返回 false，表示请求被丢弃
// 真正的漏桶算法无法提供 allow 函数（请求可以执行的判断函数），只能提供一个 push 函数，返回 true 表示请求被放到等待队列中等待执行，push 函数失败，则表明等待队列满了，请求要丢弃。
// 漏桶限流内部的实现是生产者-消费者模式，其 push 函数表示当前请求进入了请求队列，但并不表示请求可以马上被处理，这和其他限流算法的 allow 函数不同。
func (l *LeakyBucket) Push() bool {
	// 如果通道可以发送，请求被接受
	select {
	case l.queue <- struct{}{}:
		return true
	default:
		return false
	}
}

// Process 从队列中取出请求并模拟处理过程
func (l *LeakyBucket) Process() {
	for range l.queue { // 使用 range 来持续接收队列中的请求
		fmt.Println("Request processed at", time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(100 * time.Millisecond) // 模拟请求处理时间
	}
}

/*
实现原理： 通过一个固定容量的队列来模拟桶，以恒定速率从桶中取出请求进行处理，无论请求到达的频率如何，都保证请求以均匀的速度被处理，从而平滑流量并防止流量突增。

优点：
	能够强制实现固定的数据处理速率，平滑流量。
	即使面对突发流量，也能保持稳定的处理速率。

缺点：
	对于突发流量的处理不够灵活，可能会延迟处理。
	实现相对简单，但需要维护桶的状态。

漏桶算法适用于需要强制执行固定速率处理的场景，如网络流量控制、API请求限制等。通过控制令牌的添加速率，漏桶算法能够有效地避免系统因瞬时流量高峰而过载。
*/
