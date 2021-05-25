/*
要找出10000以内所有的素数，这里使用的方法是筛法，即从2开始每找到一个素数就标记所有能被该素数整除的所有数。直到没有可标记的数，剩下的就都是素数。下面以找出10以内所有素数为例，借用 CSP 方式解决这个问题。
通过4个并发处理程序得出10以内的素数表
*/

package main

import "fmt"

func Processor(seq chan int, wait chan struct{}) {
	go func() {
		prime, ok := <-seq
		if !ok {
			close(wait)
			return
		}
		fmt.Println("prime is: ", prime)
		out := make(chan int)
		Processor(out, wait)		
		for num := range seq {
			fmt.Println("num is: ", num)
			if num % prime != 0 {
				out <- num
			}				
		}
		close(out)
	}()
}

func main() {
	origin, wait := make(chan int), make(chan struct{})
	Processor(origin, wait)
	for num := 2; num < 10; num++ {
		origin <- num
	}
	close(origin)
	<-wait
}