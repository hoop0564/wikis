package cancel_by_close

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(ch chan struct{}) bool {
	// 阻塞的
	//_, ok := <-ch
	//return ok //关闭的channel还可以读！但是读回的ok值为false！
	//return ok == false

	// 非阻塞的
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {

	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 200)
				fmt.Println(i)
			}
			fmt.Println(i, "cancelled")
		}(i, cancelChan)
	}

	time.Sleep(time.Millisecond * 1000)
	close(cancelChan)
	time.Sleep(time.Millisecond * 1000)

}
