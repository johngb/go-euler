package main //same package name as source file

import (
	"testing" //import go package for testing related functionality
)

func Benchmark_Func1(b *testing.B) { // Form: Benchmark_functionName
	for i := 0; i < b.N; i++ {
		euler1(100)
	}
}

func Benchmark_Func2(b *testing.B) { // Form: Benchmark_functionName
	for i := 0; i < b.N; i++ {
		euler2(100)
	}
}
