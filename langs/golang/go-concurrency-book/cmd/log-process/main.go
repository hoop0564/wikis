package main

import (
	"fmt"
	"strings"
	"time"
)

// 从文件夹读取日志，分析后，存入influxDB

type Reader interface {
	Read(rc chan string)
}

type Writer interface {
	Write(wc chan string)
}

type LogProcess struct {
	rc    chan string // 读取chan
	wc    chan string // 写入chan
	read  Reader
	write Writer
}

type ReadFromFile struct {
	path string // 读取文件的路径
}

func (r *ReadFromFile) Read(rc chan string) {
	// 读取模块
	line := "message"
	rc <- line
}

// Process 解析模块
func (l *LogProcess) Process() {
	data := <-l.rc
	l.wc <- strings.ToUpper(data)
}

// WriteToInfluxDB 写入模块
type WriteToInfluxDB struct {
	influxDBDsn string // info data source
}

func (w *WriteToInfluxDB) Write(wc chan string) {
	fmt.Println(<-wc)
}

func main() {

	r := &ReadFromFile{
		path: "/tmp/access.log",
	}
	w := &WriteToInfluxDB{
		influxDBDsn: "username&password&host:port",
	}
	lp := &LogProcess{
		rc:    make(chan string),
		wc:    make(chan string),
		read:  r,
		write: w,
	}

	// 此写法也可以！但golang考虑到可读性，用下面的写法也OK
	// go (*lp).ReadFromFile()
	go lp.read.Read(lp.rc)
	go lp.Process()
	go lp.write.Write(lp.wc)

	time.Sleep(1 * time.Second)
	fmt.Println("exit")
}
