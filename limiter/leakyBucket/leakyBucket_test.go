package leakyBucket

import (
	"fmt"
	"testing"
	"time"
)

func TestLeakBucket(t *testing.T) {
	// 创建一个容量为5的漏桶
	limiter := NewLeakBucket(2)

	// 启动请求处理循环
	go limiter.Process()

	// 模拟请求
	for i := 0; i < 10; i++ {
		accepted := limiter.Push()
		if accepted {
			fmt.Printf("Request %d accepted at %v\n", i+1, time.Now().Format("2006-01-02 15:04:05"))
		} else {
			fmt.Printf("Request %d rejected at %v\n", i+1, time.Now().Format("2006-01-02 15:04:05"))
		}
	}
	time.Sleep(2 * time.Second)
}
