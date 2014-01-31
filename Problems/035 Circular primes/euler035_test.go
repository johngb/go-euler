package main //same package name as source file

import (
	"testing" //import go package for testing related functionality
)

func Benchmark_countCircularPrimes1(b *testing.B) { // Form: Benchmark_functionName
	for i := 0; i < b.N; i++ {
		// Funcion call to function to benchmark
		countCircularPrimes1(1000000)
	}
}

func Benchmark_countCircularPrimes2(b *testing.B) { // Form: Benchmark_functionName
	for i := 0; i < b.N; i++ {
		// Funcion call to function to benchmark
		countCircularPrimes2(1000000)
	}
}
func Benchmark_countCircularPrimes3(b *testing.B) { // Form: Benchmark_functionName
	for i := 0; i < b.N; i++ {
		// Funcion call to function to benchmark
		countCircularPrimes3(1000000)
	}
}
