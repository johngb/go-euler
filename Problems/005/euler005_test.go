package main //same package name as source file

import (
	"testing" //import go package for testing related functionality
)

func Benchmark_smallestNum1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		smallestNum1(20)
	}
}

func Benchmark_smallestNum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		smallestNum2(20)
	}
}
