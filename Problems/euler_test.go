package main //same package name as source file

import (
	"testing" //import go package for testing related functionality
)

func Benchmark_(b *testing.B) { // Form: Benchmark_functionName
	for i := 0; i < b.N; i++ {
		// Funcion call to function to benchmark
	}
}
