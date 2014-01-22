/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we
can see that the 6th prime is 13.

What is the 10 001st prime number?
*/

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

// Uses sieve of Eratosthenes, so faster than brute force.
func nthPrime1(num int) int {
	// defer timeTrack(time.Now(), "nthPrime1()") // Timer function

	// initialise an array of with num elements
	// TODO: Lots of optimisation here
	// sieveSize determines the maximum size of the prime number
	// but if it's too big, it makes the program inefficient
	const sieveSize = 1000000
	sieve := make([]bool, sieveSize+2)
	// intialise slice to contain prime numbers
	primeCounter := 0

	// index = 2
	// while i <= num
	for i := 2; i <= sieveSize; i++ {

		if sieve[i] == false {
			// pf("i = %d\n", i)

			// i should be a prime number, so add to list of primes
			primeCounter += 1
			if primeCounter == num {
				// pf("i = %d, primeCounter = %d\n", i, primeCounter)
				return i
				// return primes
			}

			for j := 2; i*j <= sieveSize; j++ {
				sieve[i*j] = true
			}
		}
	}
	// should never happen
	p("error - sieveSize is too small")
	return 0
}

// Brute force approach. Slow, but more memory efficient
func nthPrime2(num int) int {
	// defer timeTrack(time.Now(), "nthPrime2()") // Timer function

	// array of prime numbers
	primes := make([]int, 0)
	// seed the slice with a prime.
	primes = append(primes, 2)
	// start counting from 3. stop when we break
	for i := 3; len(primes) < num; i++ {
		for j := 0; j < len(primes); j++ {

			if i%primes[j] == 0 {
				// i is not a prime
				break

			} else if j == len(primes)-1 {
				// is a prime, so add to list
				primes = append(primes, i)
				// p(i)
			}
		}

	}
	return primes[num-1]
}

// TODO: Make a better sieve.
// try follow http://en.wikibooks.org/wiki/Efficient_Prime_Number_Generating_Algorithms
func primeSieve1(num int) []int {
	defer timeTrack(time.Now(), "primeSieve1()") // Timer function

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

func main() {
	// p(nthPrime1(10001))
	// p(nthPrime2(10001))
	st := "0123456789"
	st1 := st[0]
	p(st1)

}
