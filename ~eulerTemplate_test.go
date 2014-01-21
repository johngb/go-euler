package main //same package name as source file

import (
	"testing" //import go package for testing related functionality
)

func Benchmark_Func1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Body here
	}
}
