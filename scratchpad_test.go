package main

import (
	"testing"
)

func Benchmark_withCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withCap()
	}
}

func Benchmark_noCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		noCap()
	}
}
