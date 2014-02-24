package main //same package name as source file

import (
	"testing" //import go package for testing related functionality
)

func Benchmark_collatzWithMaxLength1(b *testing.B) { // Form: Benchmark_functionName
	for i := 0; i < b.N; i++ {
		collatzWithMaxLength1(1e6)
	}
}

func Benchmark_collatzWithMaxLength2(b *testing.B) { // Form: Benchmark_functionName
	for i := 0; i < b.N; i++ {
		collatzWithMaxLength2(1e6)
	}
}
