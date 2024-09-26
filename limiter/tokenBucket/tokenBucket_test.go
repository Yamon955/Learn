package tokenBucket

import (
	"fmt"
	"testing"
	"time"
)

func TestNewTokenBucket(t *testing.T) {
	// 创建一个新的令牌桶实例，桶的容量为10，每秒填充2个令牌。
	limiter := NewTokenBucket(10, 2)

	// 模拟请求，观察限流效果。
	// 循环15次，每次请求判断是否被允许。
	for i := 0; i < 15; i++ {
		if limiter.Allow() {
			fmt.Println("Request", i+1, "allowed")
		} else {
			fmt.Println("Request", i+1, "rejected")
		}
	}

	// 创建一个新的令牌桶实例，桶的容量为10，每秒填充2个令牌。
	limiter2 := NewTokenBucket(10, 2)
	for i := 0; i < 15; i++ {
		if limiter2.Allow() {
			fmt.Println("Request", i+1, "allowed")
		} else {
			fmt.Println("Request", i+1, "rejected")
		}
		time.Sleep(500 * time.Millisecond) // 确保每两个请求之后重新添加2个令牌，使每个请求全部被处理
	}
}
