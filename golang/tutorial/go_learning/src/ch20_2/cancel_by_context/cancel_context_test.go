package cancel_by_context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 200)
				fmt.Println(i)
			}
			fmt.Println(i, "cancelled")
		}(i, ctx)
	}

	time.Sleep(time.Millisecond * 1000)
	cancel()
	time.Sleep(time.Millisecond * 1000)

}
