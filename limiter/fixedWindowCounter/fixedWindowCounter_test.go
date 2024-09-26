package fixedWindowCounter

import (
	"fmt"
	"testing"
	"time"
)

func TestNewFixedWindowCounter(t *testing.T) {
	limiter := NewFixedWindowCounter(10, time.Minute)
	// 模拟15个请求，观察限流效果。
	for i := 0; i < 15; i++ {
		if limiter.Allow() {
			fmt.Println("Request", i+1, "allowed")
		} else {
			fmt.Println("Request", i+1, "rejected")
		}
	}
}
