// Package cmd 提供命令字解析与执行功能。
package cmd

import (
	"flag"
	"sort"
)

// CMDs 列举了支持的命令字。
var CMDs = cmds{}

// Command 提供命令字执行接口。
type Command interface {
	// Get 获取命令解析器。
	Get() (_ *flag.FlagSet, needArgs bool)
	// Process 执行命令字。
	Process() string
}

type cmds map[string]Command

// register 提供了在各命令字 init 函数中注册该命令字的功能, 命令字不能冲突, 否则 panic。
func register(n string, c Command) {
	if _, ok := CMDs[n]; ok {
		panic("conflict cmd")
	}
	CMDs[n] = c
}

// GetKeys 返回所支持的命令字列表。
func (c cmds) GetKeys() []string {
	var s []string
	for k := range c {
		s = append(s, k)
	}
	sort.Strings(s)
	return s
}
