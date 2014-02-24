package main //same package name as source file

import (
	"testing" //import go package for testing related functionality
)

func Benchmark_Pytrip1(b *testing.B) { // Form: Benchmark_functionName
	for i := 0; i < b.N; i++ {
		PyTrip1(1000)
	}
}

func Benchmark_Pytrip2(b *testing.B) { // Form: Benchmark_functionName
	for i := 0; i < b.N; i++ {
		PyTrip2(1000)
	}
}
