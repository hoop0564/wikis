package gc_friendly

import "testing"

const (
	numOfElems = 100000
	times      = 1000
)

func TestAutoGrow(t *testing.T) {
	for i := 0; i < times; i++ {
		//s := []int{}
		var s []int
		for j := 0; j < numOfElems; j++ {
			s = append(s, j)
		}
	}
}

func TestProperInit(t *testing.T) {
	for i := 0; i < times; i++ {
		s := make([]int, 0, 100000)
		for j := 0; j < numOfElems; j++ {
			s = append(s, j)
		}
	}
}

func TestOverSizeInit(t *testing.T) {
	for i := 0; i < times; i++ {
		s := make([]int, 0, 800000)
		for j := 0; j < numOfElems; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkAutoGrow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//s := []int{}
		var s []int
		for j := 0; j < numOfElems; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkProperInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, 100000)
		for j := 0; j < numOfElems; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkOverSizeInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, 800000)
		for j := 0; j < numOfElems; j++ {
			s = append(s, j)
		}
	}
}

/*
=== RUN   TestAutoGrow
--- PASS: TestAutoGrow (0.34s)
=== RUN   TestProperInit
--- PASS: TestProperInit (0.11s)
=== RUN   TestOverSizeInit
--- PASS: TestOverSizeInit (0.31s)
goos: darwin
goarch: amd64
pkg: geek.time.learn/src/ch49_gc/gc_friendly
BenchmarkAutoGrow-12                3577            321697 ns/op
BenchmarkProperInit-12              9421            115572 ns/op
BenchmarkOverSizeInit-12            2965            348425 ns/op
*/
