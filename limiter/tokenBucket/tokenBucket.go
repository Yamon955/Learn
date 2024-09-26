package tokenBucket

import (
	"sync"
	"time"
)

// TokenBucket 结构体实现令牌桶限流算法。
// - mu 用于同步访问，保证并发安全。
// - capacity 定义桶的容量，即桶中最多可以存放的令牌数。
// - tokens 表示桶中当前的令牌数。
// - refillRate 是令牌的填充速率，表示每秒向桶中添加的令牌数。
// - lastRefill 记录上次填充令牌的时间。
type TokenBucket struct {
	mux        sync.Mutex
	capacity   int
	tokens     int
	refillRate float64
	lastRefill time.Time
}

// NewTokenBucket 构造函数初始化 TokenBucket 实例。
// - capacity 参数定义了桶的容量。
// - refillRate 参数定义了每秒向桶中添加的令牌数。
func NewTokenBucket(capacity int, refillRate float64) *TokenBucket {
	// 初始化时桶被填满，tokens 和 capacity 相等。
	// lastRefill 设置为当前时间。
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

// Allow 方法用于判断当前请求是否被允许。
func (t *TokenBucket) Allow() bool {
	t.mux.Lock()
	defer t.mux.Unlock()

	now := time.Now() // 获取当前时间。

	// 计算自上次填充以来经过的秒数，并转换为float64类型。
	timeElapsed := float64(now.Unix() - t.lastRefill.Unix())

	// 根据 refillRate 计算应该添加的令牌数。
	tokensToAdd := t.refillRate * timeElapsed

	// 更新令牌数，但不超过桶的容量。
	t.tokens += int(tokensToAdd)
	if t.tokens > t.capacity {
		t.tokens = t.capacity // 确保令牌数不超过桶的容量。
	}

	// 如果桶中有令牌，则移除一个令牌并允许请求通过。
	if t.tokens > 0 {
		t.tokens--         // 移除一个令牌。
		t.lastRefill = now // 更新上次填充时间到当前时间。
		return true
	}

	// 如果桶中无令牌，则请求被拒绝。
	return false
}

/*
实现原理： 令牌桶算法使用一个令牌桶来调节数据流的速率，允许一定程度的流量突发。桶初始时为空，并以固定的速率填充令牌，直至达到预设的容量上限。与漏桶算法不同，令牌桶算法在桶未满时，可以在每个时间间隔内向桶中添加多个令牌，从而积累处理突发请求的能力。当请求到达时，如果桶中存在令牌，算法会从桶中移除相应数量的令牌来处理请求。如果桶中的令牌不足，请求将被延迟处理或根据策略拒绝服务。如果桶已满，额外的令牌将不会被添加，确保了令牌数量不会超过桶的容量限制。

优点：
	允许一定程度的突发流量，更加灵活。
	可以平滑流量，同时在桶未满时快速处理请求。

缺点：
	实现相对复杂，需要维护桶的状态和时间。
	对于计算和同步的要求更高。

令牌桶算法适用于需要处理突发流量的场景，如网络通信、API调用等。通过控制令牌的填充速率和桶的容量，令牌桶算法能够有效地平衡流量，防止系统过载，同时允许在短期内处理更多的请求。

漏桶算法通过固定速率释放令牌来控制请求处理速度，而令牌桶算法允许以固定速率积累令牌，以支持突发请求并保持平均处理速率。
漏桶和令牌最大的不同是：漏桶处理请求的速度是恒定的，后端的处理服务不会出现流量波峰。
更形象解释：漏桶算法在每个固定时间间隔向桶中添加一个令牌，而令牌桶算法则以一定的速率在每个时间间隔内向桶中添加多个令牌，直到达到桶的容量上限。
*/
