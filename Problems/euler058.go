/*
Starting with 1 and spiralling anticlockwise in the following way, a square
spiral with side length 7 is formed.

37 36 35 34 33 32 31
38 17 16 15 14 13 30
39 18  5  4  3 12 29
40 19  6  1  2 11 28
41 20  7  8  9 10 27
42 21 22 23 24 25 26
43 44 45 46 47 48 49

It is interesting to note that the odd squares lie along the bottom right
diagonal, but what is more interesting is that 8 out of the 13 numbers lying
along both diagonals are prime; that is, a ratio of 8/13 ≈ 62%.

If one complete new layer is wrapped around the spiral above, a square spiral
with side length 9 will be formed. If this process is continued, what is the
side length of the square spiral for which the ratio of primes along both
diagonals first falls below 10%?
*/

package main

import (
	"fmt"
	"math/big"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// timeTrack is used for basic benchmarking in other functions
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

func makePrimeMap(max int) map[int]bool {
	defer timeTrack(time.Now(), "makePrimeMap") // Timer function

	sieve := make([]bool, max/2+1)
	// initialise the prime list with the only even prime
	primeMap := map[int]bool{2: true}

	// 3 is the first odd prime
	for i := 3; i <= max; i += 2 {
		if sieve[i/2] == false {
			// i should be a prime number, so add to list of primes
			primeMap[i] = true
			// any odd * even = even, so avoid all even multiples
			for j := 3; i*j <= max; j += 2 {
				sieve[i*j/2] = true
			}
		}
	}
	return primeMap
}

func diagonalSumOnSpiral(lim int) int {
	defer timeTrack(time.Now(), "diagonalSumOnSpiral()")

	primeMap := makePrimeMap(lim)

	cornerPrimeCount := float64(0)
	cornerNonPrimeCount := float64(1) // the centre isn't a prime
	counter := 1
	width := 1
	for loop := 2; counter < lim; loop++ {
		// do 4 times for each loop
		inc := ((loop - 1) * 2)
		for i := 1; i <= 4; i++ {
			counter += inc
			if primeMap[counter] {
				cornerPrimeCount += 1.0
			} else {
				cornerNonPrimeCount += 1.0
			}
		}
		width += 2
		ratio := (cornerPrimeCount / (cornerNonPrimeCount + cornerPrimeCount))
		// p("ratio =", ratio)
		if ratio < 0.1 {
			p("non-primes:", cornerNonPrimeCount, " primes:", cornerPrimeCount)
			return width
		}
	}
	return int((cornerPrimeCount / (cornerNonPrimeCount + cornerPrimeCount)) * 1000)
}

// uses a probablistic prime test, rather than a deterministic one.  Much faster, but slightly less accurate.
func diagonalSumOnSpiral1(lim int) int {
	defer timeTrack(time.Now(), "diagonalSumOnSpiral1()")

	cornerPrimeCount := float64(0)
	cornerNonPrimeCount := float64(1) // the centre isn't a prime
	counter := 1
	width := 1
	for loop := 2; counter < lim; loop++ {
		// do 4 times for each loop
		inc := ((loop - 1) * 2)
		for i := 1; i <= 4; i++ {
			counter += inc
			bigCounter := big.NewInt(int64(counter))

			if bigCounter.ProbablyPrime(10) {
				cornerPrimeCount += 1.0
			} else {
				cornerNonPrimeCount += 1.0
			}
		}
		width += 2
		ratio := (cornerPrimeCount / (cornerNonPrimeCount + cornerPrimeCount))
		// p("ratio =", ratio)
		if ratio < 0.1 {
			p("non-primes:", cornerNonPrimeCount, " primes:", cornerPrimeCount)
			return width
		}
	}
	p("Not found. Ratio:")
	return int((cornerPrimeCount / (cornerNonPrimeCount + cornerPrimeCount)) * 1000)
}

func main() {
	// p(primeMap(100))
	// p(diagonalSumOnSpiral(100000000))
	p(diagonalSumOnSpiral1(1000000000))

}
