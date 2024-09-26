package slidingWindowLimiter

import (
	"fmt"
	"testing"
	"time"
)

func TestNewSlidingWindowLimiter(t *testing.T) {
	limiter := NewSlidingWindowLimiter(1, time.Second, 10*time.Millisecond)
	for i := 0; i < 100; i++ {
		time.Sleep(5 * time.Millisecond)
		if limiter.Allow() {
			fmt.Println("Request", i+1, "allowed")
		} else {
			fmt.Println("Request", i+1, "rejected")
		}
	}
}
