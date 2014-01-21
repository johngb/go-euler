package main

import (
	"fmt"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// Used to time other functions to compare speed
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

/*
The sum of the squares of the first ten natural numbers is,
12 + 22 + ... + 102 = 385

The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)2 = 552 = 3025

Hence the difference between the sum of the squares of the first
ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first
one hundred natural numbers and the square of the sum.
*/

// Brute force approach
func sumOfSquares1(n int) int {
	// defer timeTrack(time.Now(), "sumOfSquares1()") // Timer function

	// for each numer in range 1:n
	sum := 0
	for i := n; i > 0; i-- {
		// pf("%d + %d = %d\n", sum, i*i, sum+i*i)
		sum += i * i
	}
	return sum
}

// Elegant maths approach using the series
// sum of squares to n = n(n+1)(2n+1)/6
func sumOfSquares2(n int) int {
	// defer timeTrack(time.Now(), "sumOfSquares2()") // Timer function

	sum := n * (n + 1) * (2*n + 1) / 6
	return sum
}

//Brute force approach
func squareOfSum1(n int) int {
	// defer timeTrack(time.Now(), "squareOfSum1()") // Timer function

	// for each numer in range 1:n
	sum := 0
	for i := n; i > 0; i-- {
		// pf("%d + %d = %d\n", sum, i, sum+i)
		sum += i
	}
	return sum * sum
}

// Elegant maths approach using the series
// square of sums to n = (n^2 + n)^2 / 4
func squareOfSum2(n int) int {
	// defer timeTrack(time.Now(), "squareOfSum2()") // Timer function

	squareOfSum := (n*n*n*n + 2*n*n*n + n*n) / 4

	return squareOfSum
}

func euler1(n int) int {
	return squareOfSum1(n) - sumOfSquares1(n)
}

func euler2(n int) int {
	return squareOfSum2(n) - sumOfSquares2(n)
}

func main() {
	p(euler1(100))
	p(euler2(100))
}
