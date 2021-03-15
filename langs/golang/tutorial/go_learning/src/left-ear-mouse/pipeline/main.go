package main

import "fmt"

// Channel 转发函数

func echo(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			fmt.Println("echo", n)
			// 写入out，等待外部读取；out被读取后，继续for循环，直至协程完工退出
			out <- n
		}
		close(out)
	}()
	return out
}

// 平方函数
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
			fmt.Println("sq", n)
		}
		close(out)
	}()
	return out
}

// 过滤奇数函数
func odd(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if n%2 != 0 {
				out <- n
				fmt.Println("odd", n)
			}
		}
		close(out)
	}()
	return out
}

// 求和函数
func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum = 0
		for n := range in {
			sum += n
			fmt.Println("sum", n)
		}
		out <- sum
		close(out)
	}()
	return out
}

type EchoFunc func([]int) <-chan int
type PipeFunc func(<-chan int) <-chan int

func pipeline(nums []int, echoF EchoFunc, pipeFns ...PipeFunc) <-chan int {
	ch := echoF(nums)
	for i := range pipeFns {

		ch = pipeFns[i](ch)
	}
	return ch
}

func main() {

	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for n := range pipeline(nums, echo, odd, sq, sum) {
		fmt.Println("pipeline", n)
	}
}
