package main

import "testing"

var benchResult int // global sink to avoid compiler optimizations

func BenchmarkPopCount(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = PopCount(uint64(i))
	}
	benchResult = r
}

func BenchmarkPopCountWithoutLoop(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = PopCountWithoutFor(uint64(i))
	}
	benchResult = r
}
