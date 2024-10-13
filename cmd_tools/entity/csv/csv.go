// Package csv 提供线程安全CSV操作函数。
package csv

import (
	"encoding/csv"
	"io"
	"os"
	"sync"
)

type Writer struct {
	w    *csv.Writer
	lock sync.Mutex
	f    *os.File
}

// NewWriter 创建写入者
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		w: csv.NewWriter(w),
	}
}

func NewWriterWithFile(uri string) (*Writer, error) {
	file, err := os.OpenFile(uri, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &Writer{
		w: csv.NewWriter(file),
		f: file,
	}, nil
}

// Error 出错信息
func (w *Writer) Error() error {
	return w.w.Error()
}

func (w *Writer) Close() {
	w.Flush()
	if w.f != nil {
		w.f.Close()
	}
}

func (w *Writer) Flush() {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.w.Flush()
}

func (w *Writer) Write(record []string) error {
	w.lock.Lock()
	defer w.lock.Unlock()
	return w.w.Write(record)
}

func (w *Writer) WriteAll(records [][]string) error {
	w.lock.Lock()
	defer w.lock.Unlock()
	return w.w.WriteAll(records)
}
