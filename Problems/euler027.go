/*
Euler discovered the remarkable quadratic formula:

n² + n + 41

It turns out that the formula will produce 40 primes for the consecutive
values n = 0 to 39. However, when n = 40, 402 + 40 + 41 = 40(40 + 1) + 41 is
divisible by 41, and certainly when n = 41, 41² + 41 + 41 is clearly divisible
by 41.

The incredible formula  n² − 79n + 1601 was discovered, which produces 80
primes for the consecutive values n = 0 to 79. The product of the
coefficients, −79 and 1601, is −126479.

Considering quadratics of the form:

n² + an + b, where |a| < 1000 and |b| < 1000

where |n| is the modulus/absolute value of n e.g. |11| = 11 and |−4| = 4 Find
the product of the coefficients, a and b, for the quadratic expression that
produces the maximum number of primes for consecutive values of n, starting
with n = 0.
*/

// ------------------------------------------------
// JGB: could add optimisations here, as we know that b has to be a prime, as when n = 0,
// quadraticValue = b.  So restrict all b values to primes.  Possibly also generate a list of primes
// that are used to check against rather than calculating each one every time.

package main

import (
	"fmt"
	"math"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// timeTrack is used for basic benchmarking in other functions
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

func quadraticValue(a int, b int, n int) int {
	return n*n + a*n + b
}

// isPrime returns true if 'pPrime' is a prime
func isPrime(pPrime int) bool {
	if pPrime < 0 {
		return false
	}
	maxFactor := int(math.Sqrt(float64(pPrime)))
	for i := 2; i <= maxFactor; i++ {
		if pPrime%i == 0 {
			return false
		}
	}
	return true
}

// nextPrime returns the next prime number greater than 'prime'
func nextPrime(prime int) int {
	// defer timeTrack(time.Now(), "nextPrime()")

	var pPrime int
	// makes sure that the next number isn't even
	if prime%2 == 0 {
		pPrime = prime + 1
	} else {
		pPrime = prime + 2
	}
	for {
		if isPrime(pPrime) {
			break
		}
		pPrime += 2
	}
	return pPrime
}

func checkPrimeRun(a int, b int) int {
	// defer timeTrack(time.Now(), "checkPrimeRun()")

	n := 0
	for n = 0; isPrime(quadraticValue(a, b, n)); n++ {
		// pf("n = %d, quad = %d\n", n, quadraticValue(a, b, n))
	}
	return n
}

func longestPrimeRun(limits int) int {
	defer timeTrack(time.Now(), "longestPrimeRun()")
	var a int
	var b int
	bestCoefficients := []int{-limits - 1, limits + 1}
	bestRun := 0
	run := 0
	for a = -limits; a <= limits; a++ {
		for b = -limits; b <= limits; b++ {
			// check to see if b is even, as b has to be a prime
			if b%2 == 0 && b != 2 {
				continue
			}
			run = checkPrimeRun(a, b)
			if run > bestRun {
				bestRun = run
				bestCoefficients[0] = a
				bestCoefficients[1] = b
			}
		}
	}
	pf("a = %d, b = %d, bestrun = %d\n", bestCoefficients[0], bestCoefficients[1], bestRun)
	return bestCoefficients[0] * bestCoefficients[1]
}

func main() {
	// p(isPrime(1681))
	// p(checkPrimeRun(1, 41))
	p(longestPrimeRun(1000))
}
