package main //same package name as source file

import (
	"testing" //import go package for testing related functionality
)

func Benchmark_nthPrime1(b *testing.B) { // Form: Benchmark_functionName
	for i := 0; i < b.N; i++ {
		nthPrime1(10001)
	}
}

func Benchmark_nthPrime2(b *testing.B) { // Form: Benchmark_functionName
	for i := 0; i < b.N; i++ {
		nthPrime2(10001)
	}
}
