/*
The prime 41, can be written as the sum of six consecutive primes:

41 = 2 + 3 + 5 + 7 + 11 + 13

This is the longest sum of consecutive primes that adds to a prime below one-
hundred.

The longest sum of consecutive primes below one-thousand that adds to a prime,
contains 21 terms, and is equal to 953.

Which prime, below one-million, can be written as the sum of the most
consecutive primes?
*/

package main

import (
	"fmt"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// timeTrack is used for basic benchmarking in other functions
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

func primeListAndMap(max int) ([]int, map[int]bool) {

	sieve := make([]bool, max/2+1)
	// initialise the prime list with the only even prime
	primeList := []int{2}
	primeMap := map[int]bool{2: true}

	// 3 is the first odd prime
	for i := 3; i <= max; i += 2 {
		if sieve[i/2] == false {
			// i should be a prime number, so add to list of primes
			primeList = append(primeList, i)
			primeMap[i] = true
			// any odd * even = even, so avoid all even multiples
			for j := 3; i*j <= max; j += 2 {
				sieve[i*j/2] = true
			}
		}
	}
	return primeList, primeMap
}

func findLongestPrimeSum(lim int) int {
	defer timeTrack(time.Now(), "findLongestPrimeSum()")

	// find all primes < limit
	// JGB: can find a lower number, but not sure how much lower
	primes, primeMap := primeListAndMap(lim)

	longestSequence := 0
	longestPrime := 0
	sum := 0
	newsum := 0
	sequence := 0

	// vary the starting prime
	for j := 0; j < 100; j++ {
		// count from the jth prime
		for i := j; newsum < lim; i++ {
			sum = newsum
			newsum += primes[i]
			sequence += 1

			if primeMap[sum] == true {
				if sequence > longestSequence {
					longestSequence = sequence
					longestPrime = sum
				}
			}
		}
		sum = 0
		newsum = 0
		sequence = 0
	}
	p(longestSequence)
	return longestPrime
}

func main() {
	// const n = 100
	// p(primeSieve(n))
	// p(primeSieve1(n))

	p(findLongestPrimeSum(1000000))

}
