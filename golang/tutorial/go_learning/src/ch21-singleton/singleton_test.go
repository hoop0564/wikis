package ch21_singleton

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var singletonInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("create obj")
		singletonInstance = new(Singleton)
	})
	return singletonInstance
}

func TestSingleton_test(t *testing.T) {

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%d\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
