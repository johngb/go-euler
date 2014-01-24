package main //same package name as source file

import (
	"testing" //import go package for testing related functionality
)

func Benchmark_isPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome(12345654321)
	}
}

func Benchmark_palindrome1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		palindrome1()
	}
}

func Benchmark_palindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		palindrome2()
	}
}

func Benchmark_palindrome3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		palindrome3()
	}
}
