package main

import (
	"fmt"
	// "math"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// Used to time other functions to compare speed
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

// returns a list of all prime numbers <= num
func eratosthenes(num int) []int {
	// initialise an array of with num elements
	sieve := make([]bool, num+2)
	// intialise slice to contain prime numbers
	primes := make([]int, 0)

	// index = 2
	// while i <= num
	// for i := 2; i <= int(math.Sqrt(float64(num)))+1; i++ {
	for i := 2; i <= num; i++ {

		if sieve[i] == false {

			// i should be a prime number, so add to list of primes
			primes = append(primes, i)

			for j := 2; i*j <= num; j++ {
				sieve[i*j] = true
			}
		}
	}

	// return the list of primes
	return primes
}

func intPow(x, y int) int {
	if y == 1 {
		return x
	}
	workingPower := x
	for i := 2; i <= y; i++ {
		workingPower *= x
	}
	return workingPower
}

/*
2520 is the smallest number that can be divided by each of the numbers
from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the
numbers from 1 to 20?
*/

// Brute force approach.  Not efficient
func smallestNum1(n int) int {
	// defer timeTrack(time.Now(), "smallestNum1()") // Timer function

	var number int
	// start with number = n, and keep incrementing it
	for number = n; ; number++ {
		// pf("number = %d\n", number)

		// checking for each number from n down to 2
		for i := n; i > 1; i-- {
			// pf("number = %d, i = %d\n", number, i)
			// if number is not perfectly divisible by i
			if number%i != 0 {
				break
			}
			// if i = 2, we have a solution
			if i == 2 {
				return number
			}
		}
	}
	return 0 // should never happen
}

// Intelligent but more complicated approach
// Method: consider all prime numbers <= n
// multiply the highest powers of those primes that are < n
func smallestNum2(n int) int {
	// defer timeTrack(time.Now(), "smallestNum2()") // Timer function

	// find the list of prime numbers <= n
	// initialise an empty slice
	factors := make([]int, 0)
	factors = append(factors, eratosthenes(n)...)

	// for each factor
	for index := 0; index < len(factors); index++ {
		// determine how many times it fits in n
		for x := 1; ; x++ {

			// TODO: make compare to n/x for optimisation
			if intPow(factors[index], x) > n {
				factors[index] = intPow(factors[index], (x - 1))
				break
			}
		}
	}
	answer := 1
	for i := 0; i < len(factors); i++ {
		answer *= factors[i]
	}
	return answer
}

func main() {
	const x = 50
	// p(smallestNum1(x))
	p(smallestNum2(x))
	// p(intPow(2, 12))
}
