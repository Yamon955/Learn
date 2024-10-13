// Package counter 统计器
package counter

import (
	"fmt"
	"sync/atomic"
)

// Data 统计数据
type Data struct {
	Count uint64 // vid计数
	Size  uint64 // size计数 Byte
}

// Add 增加计数
func (d *Data) Add(size uint64) {
	atomic.AddUint64(&d.Count, 1)
	atomic.AddUint64(&d.Size, size)
}

// String 打印
func (d *Data) String() string {
	return fmt.Sprintf("count:%v, size:%v MB", d.Count, d.Size/1024/1024)
}
