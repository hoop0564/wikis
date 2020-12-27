package ch39_unsafe_programming

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

// 不适合的场景：无意义的强制类型转换，其值可能有丢失，比如float64和int的unsafe.Pointer转换
func TestUnsafe(t *testing.T) {
	i := 10
	f := *(*float64)(unsafe.Pointer(&i))
	t.Log(unsafe.Pointer(&i)) // 输出指针 0xc00008e1a8
	t.Log(f)                  // 输出f的值 5e-323 != 10
}

// 自定义了一个比如 type MyInt int，后面需要对int类的变量做使用MyInt的方法操作
type MyInt int

// 合理的类型转换
func TestConvert(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := *(*[]MyInt)(unsafe.Pointer(&a))
	t.Log(b) // [1 2 3 4]
}

// 合理的原子类型操作
func TestAtomic(t *testing.T) {
	var shareBufPtr unsafe.Pointer
	writeDataFn := func() {
		var data []int
		for i := 0; i < 10; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&shareBufPtr, unsafe.Pointer(&data))
	}

	readDataFn := func() {
		data := atomic.LoadPointer(&shareBufPtr)
		//t.Log(data, *(*[]int)(data))
		fmt.Println(data, *(*[]int)(data)) //输出地址和内容
	}

	var wg sync.WaitGroup
	writeDataFn()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				writeDataFn()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				readDataFn()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
