package main //same package name as source file

import (
	"testing" //import go package for testing related functionality
)

// func Benchmark_sumOfPrimes1(b *testing.B) { // Form: Benchmark_functionName
// 	for i := 0; i < b.N; i++ {
// 		sumOfPrimes1(10000)
// 	}
// }

// func Benchmark_sumOfPrimes2(b *testing.B) { // Form: Benchmark_functionName
// 	for i := 0; i < b.N; i++ {
// 		sumOfPrimes2(10000)
// 	}
// }

func Benchmark_sumOfPrimes3(b *testing.B) { // Form: Benchmark_functionName
	for i := 0; i < b.N; i++ {
		sumOfPrimes3(2000000)
	}
}
