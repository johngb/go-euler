package main //same package name as source file

import (
	"testing" //import go package for testing related functionality
)

const n = 4e6 // limit the

func Benchmark_Func1(b *testing.B) { //benchmark function starts with "Benchmark" and takes a pointer to type testing.B
	for i := 0; i < b.N; i++ { //use b.N for looping
		countFibEven1(n)
	}
}

func Benchmark_Func2(b *testing.B) { //benchmark function starts with "Benchmark" and takes a pointer to type testing.B
	for i := 0; i < b.N; i++ { //use b.N for looping
		countFibEven2(n)
	}
}

func Benchmark_Func3(b *testing.B) { //benchmark function starts with "Benchmark" and takes a pointer to type testing.B
	for i := 0; i < b.N; i++ { //use b.N for lopoing
		countFibEven3(n)
	}
}

func Benchmark_Func4(b *testing.B) { //benchmark function starts with "Benchmark" and takes a pointer to type testing.B
	for i := 0; i < b.N; i++ { //use b.N for looping
		countFibEven4(n)
	}
}

func Benchmark_Func5(b *testing.B) { //benchmark function starts with "Benchmark" and takes a pointer to type testing.B
	for i := 0; i < b.N; i++ { //use b.N for looping
		countFibEven5(n)
	}
}

func Benchmark_Func6(b *testing.B) { //benchmark function starts with "Benchmark" and takes a pointer to type testing.B
	for i := 0; i < b.N; i++ { //use b.N for looping
		countFibEven6(n)
	}
}

func Benchmark_Func7(b *testing.B) { //benchmark function starts with "Benchmark" and takes a pointer to type testing.B
	for i := 0; i < b.N; i++ { //use b.N for looping
		countFibEven7(n)
	}
}

func Benchmark_Func8(b *testing.B) { //benchmark function starts with "Benchmark" and takes a pointer to type testing.B
	for i := 0; i < b.N; i++ { //use b.N for looping
		countFibEven8(n)
	}
}
