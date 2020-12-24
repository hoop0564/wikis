package channel_mutex

import (
	"fmt"
	"testing"
	"time"
)

type Empty interface{}    // 空接口
type semaphore chan Empty // 信号量

// 信号量的P操作，获取资源
func (s *semaphore) P(n int) {

	e := new(Empty)
	for i := 0; i < n; i++ {
		*s <- e
	}
}

// 信号量的V操作，释放资源
func (s *semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-*s
	}
}

// 互斥锁的上锁
func (s *semaphore) Lock() {
	s.P(1)
}

// 互斥锁的解锁
func (s *semaphore) Unlock() {
	s.V(1)
}

// 互斥锁的使用
func printInt(s *semaphore) {
	s.Lock()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	s.Unlock()
}

func TestMutex(t *testing.T) {
	sem := make(semaphore, 1)
	go printInt(&sem)
	go printInt(&sem)

	time.Sleep(10e2)
}
