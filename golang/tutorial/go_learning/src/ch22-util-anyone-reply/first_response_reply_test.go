package ch22_util_anyone_reply

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// 仅需任一任务完成
// 比如从百度、必应、谷歌同时搜索关键词，只要有一个先回来即可

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("before goroutine num:", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(100 * time.Millisecond)
	t.Log("after goroutine num:", runtime.NumGoroutine())
}
