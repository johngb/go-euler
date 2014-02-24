/*
The number 3797 has an interesting property. Being prime itself, it is
possible to continuously remove digits from left to right, and remain prime at
each stage: 3797, 797, 97, and 7. Similarly we can work from right to left:
3797, 379, 37, and 3.

Find the sum of the only eleven primes that are both truncatable from left to
right and right to left.

NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
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

func mapOfPrimes(min, max int) map[int]bool {
	// defer timeTrack(time.Now(), "mapOfPrimes") // Timer function

	sieve := make([]bool, max+2)
	primeMap := make(map[int]bool)

	// 2 is the first prime
	for i := 2; i <= max; i++ {
		if sieve[i] == false {
			// i should be a prime number, so add to list of primes
			primeMap[i] = true
			for j := 2; i*j <= max; j++ {
				sieve[i*j] = true
			}
		}
	}
	return primeMap
}

// truncate returns the sum of all truncateable primes below 'lim'
func truncate(lim int) int {
	defer timeTrack(time.Now(), "truncate()")

	primeMap := mapOfPrimes(2, lim)
	sum, count := 0, 0

outerloop:
	for n := 10; n < lim; n++ {

		if primeMap[n] == true {

			digitCounter := 0
			tempN := n
			for {
				// take a digit off the right
				tempN /= 10
				if tempN == 0 {
					break
				}
				digitCounter++
				if primeMap[tempN] == false {
					continue outerloop
				}
			}
			div := 1
			for ; digitCounter > 0; digitCounter-- {
				div *= 10
			}
			for div > 1 {
				tempN = n - (n/div)*div
				div /= 10
				if primeMap[tempN] == false {
					continue outerloop
				}
			}
			p(n)
			sum += n
			count += 1
		}
	}
	p(count)
	return sum
}

func main() {
	truncate(1000000)

}
