package main

import (
	"fmt"
	"strings"
	"time"
)

// 从文件夹读取日志，分析后，存入influxDB

type LogProcess struct {
	rc          chan string // 读取chan
	wc          chan string // 写入chan
	path        string      // 读取文件的路径
	influxDBDsn string
}

// ReadFromFile 读取模块
func (l *LogProcess) ReadFromFile() {

	line := "message"
	l.rc <- line
}

// Process 解析模块
func (l *LogProcess) Process() {
	data := <-l.rc
	l.wc <- strings.ToUpper(data)
}

// WriteToInfluxDB 写入模块
func (l *LogProcess) WriteToInfluxDB() {
	fmt.Println(<-l.wc)
}

func main() {

	lp := &LogProcess{
		rc: make(chan string),
		wc: make(chan string),
		path:        "/tmp/access.log",
		influxDBDsn: "username&password@host:port",
	}

	// 此写法也可以！但golang考虑到可读性，用下面的写法也OK
	// go (*lp).ReadFromFile()
	go lp.ReadFromFile()
	go lp.Process()
	go lp.WriteToInfluxDB()

	time.Sleep(1 * time.Second)
	fmt.Println("exit")
}
