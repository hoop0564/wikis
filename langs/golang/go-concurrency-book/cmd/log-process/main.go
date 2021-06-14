package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// 从文件夹读取日志，分析后，存入influxDB

type Reader interface {
	Read(rc chan []byte)
}

type Writer interface {
	Write(wc chan string)
}

type LogProcess struct {
	rc    chan []byte // 读取chan
	wc    chan string // 写入chan
	read  Reader
	write Writer
}

type ReadFromFile struct {
	path string // 读取文件的路径
}

func (r *ReadFromFile) Read(rc chan []byte) {
	// 读取模块
	f, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("open file error: %s", err.Error()))
	}

	// 从文件末尾开始逐行读取文件内容
	f.Seek(0, 2)
	rd := bufio.NewReader(f)

	for {
		line, err := rd.ReadBytes('\n')
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("ReadBytes error: %s", err.Error()))
		}
		rc <- line[:len(line)-1]
	}
}

// Process 解析模块
func (l *LogProcess) Process() {
	for v := range l.rc {
		l.wc <- strings.ToUpper(string(v))
	}
}

// WriteToInfluxDB 写入模块
type WriteToInfluxDB struct {
	influxDBDsn string // info data source
}

func (w *WriteToInfluxDB) Write(wc chan string) {
	for v := range wc {
		fmt.Println(v)
	}
}

func main() {

	r := &ReadFromFile{
		path: "access.log",
	}
	w := &WriteToInfluxDB{
		influxDBDsn: "username&password&host:port",
	}
	lp := &LogProcess{
		rc:    make(chan []byte),
		wc:    make(chan string),
		read:  r,
		write: w,
	}

	// 此写法也可以！但golang考虑到可读性，用下面的写法也OK
	// go (*lp).ReadFromFile()
	go lp.read.Read(lp.rc)
	go lp.Process()
	go lp.write.Write(lp.wc)

	time.Sleep(100 * time.Second)
	fmt.Println("exit")
}
