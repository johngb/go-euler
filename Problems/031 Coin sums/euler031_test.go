package main //same package name as source file

import (
	"testing" //import go package for testing related functionality
)

func Benchmark_numberOfWaysToMake2Pounds(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numberOfWaysToMake2Pounds()
	}
}

func Benchmark_numberOfWaysToMakeX1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numberOfWaysToMakeX1(200)
	}
}

func Benchmark_numberOfWaysToMakeX2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numberOfWaysToMakeX2(200)
	}
}
