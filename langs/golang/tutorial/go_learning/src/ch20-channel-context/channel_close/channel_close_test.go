package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

// 消息生产者
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

// 消息消费者
func dataConsumer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			data, ok := <-ch
			if ok {
				fmt.Println(data)
			} else {
				fmt.Println("close")
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)

	wg.Add(1)
	dataConsumer(ch, &wg)

	wg.Add(1)
	dataConsumer(ch, &wg)

	wg.Wait()
}
