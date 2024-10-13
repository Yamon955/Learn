package cmd

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"golang.org/x/time/rate"

	"github.com/Yamon955/Learn/cmd_tools/entity/batch"
	"github.com/Yamon955/Learn/cmd_tools/entity/counter"
	mycsv "github.com/Yamon955/Learn/cmd_tools/entity/csv"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

func init() {
	names := []string{
		"csv_traversal",
	}
	for _, name := range names {
		register(name, newCSVTraversal(name))
	}
}

// CSVTraversal 遍历csv文件，可指定筛选列col，以及筛选条件condition, 将符合条件的vid输出到 outputfile 中
type CSVTraversal struct {
	name       string
	cmd        *flag.FlagSet
	file       *string // 遍历单个文件
	path       *string // 遍历文件夹下所有文件
	qps        *int    // qps控制并发速度
	task       *int    // 协程数
	index      *int    // 从指定索引位置开始遍历
	col        *int    // 要筛选的列
	condition  *string // 筛选条件
	outputFile *string // 输出的文件名
}

// CSVTraversalInfo vid信息
type CSVTraversalInfo struct {
	Vid   string `csv:"vid"`   // video ID
	Genre string `csv:"genre"` // 电影类型
}

func newCSVTraversal(name string) *CSVTraversal {
	cmd := flag.NewFlagSet(name, flag.ExitOnError)
	return &CSVTraversal{
		name:       name,
		cmd:        cmd,
		qps:        cmd.Int("qps", 500, "QPS"),
		task:       cmd.Int("task", 1000, "task number"),
		file:       cmd.String("file", "video.csv", "input video csv file"),
		outputFile: cmd.String("output", "output.csv", "output video csv file"),
		index:      cmd.Int("index", 0, "start index"),
		path:       cmd.String("path", "", "input file path"),
		col:        cmd.Int("col", 0, "column number"),
	}
}

// Get 获取命令解析器。
func (t *CSVTraversal) Get() (*flag.FlagSet, bool) {
	return t.cmd, true
}

// Process 执行命令字。
func (t *CSVTraversal) Process() string {
	var (
		ctx       = trpc.BackgroundContext()
		limiter   = rate.NewLimiter(rate.Limit(*t.qps), 1)
		group     = batch.New(*t.task)
		count     = 0
		nextIndex = *t.index
		counter   = &counter.Data{}
	)
	log.Infof("start load file:%v", *t.file)
	file, err := os.OpenFile(*t.file, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err.Error()
	}
	reader := bufio.NewReader(file)
	csvReader := csv.NewReader(reader)
	csvwriter, err := mycsv.NewWriterWithFile(*t.outputFile)
	if err != nil {
		return err.Error()
	}
	defer func() {
		file.Close()
		csvwriter.Close()
	}()
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		log.Infof("%v", record)
		count++
		fmt.Println(count)
		if count <= nextIndex {
			continue
		}
		if len(record) < 1 {
			log.ErrorContextf(ctx, "traverse vid record:%+v", record)
			continue
		}
		vi := &CSVTraversalInfo{
			Vid: record[0],
		}
		for !limiter.Allow() {
			continue
		}
		if count%2 == 0 {
			log.InfoContextf(ctx, "traverse video count:%v, vid:%v", count, vi.Vid)
		}
		group.Add()
		go func(vi *CSVTraversalInfo) {
			defer group.Done()
			// 没有指定condition
			if t.condition == nil || *t.condition == "" {
				// 没有指定condition，则遍历所有
				csvwriter.Write([]string{vi.Vid, vi.Genre})
				counter.Add(0)
				return
			}
			// 不符合条件
			if vi.Genre != *t.condition {
				return
			}
			// 符合条件的电影输出到文件中
			csvwriter.Write([]string{vi.Vid, vi.Genre})
			counter.Add(0)
			return
		}(vi)
	}
	group.Wait()
	log.Infof("find vid count:%d", counter.Count)
	time.Sleep(time.Second)
	return "success"
}

//  ./main csv_traversal -file=./data/video.csv -col=4
