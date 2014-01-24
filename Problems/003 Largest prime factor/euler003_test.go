// package euler003 //same package name as source file
package main

import (
	"testing" //import go package for testing related functionality
)

// const n = 20000 // constant to test the functions with

func Benchmark_eratosthenes1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		eratosthenes1(800)
	}
}

func Benchmark_eratosthenes2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		eratosthenes2(800)
	}
}

func Benchmark_eratosthenes3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		eratosthenes3(800)
	}
}

func Benchmark_biggestPrimeFactor1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		biggestPrimeFactor1(520000)
	}
}

func Benchmark_biggestPrimeFactor2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		biggestPrimeFactor2(600851475143)
	}
}

func Benchmark_biggestPrimeFactor3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		biggestPrimeFactor3(600851475143)
	}
}

func Benchmark_factors1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factors1(600851475143)
	}
}

func Benchmark_factors2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factors2(600851475143)
	}
}

func Benchmark_factors3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factors3(600851475143)
	}
}

func Benchmark_factors4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factors4(600851475143)
	}
}
