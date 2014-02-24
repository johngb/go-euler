/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

package main

import (
	"fmt"
	"math"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// Used to time other functions to compare speed
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

// Basic brute force
func sumOfPrimes1(lengthToSum int) int {
	// defer timeTrack(time.Now(), "sumOfPrimes1()") // Timer function

	primes := []int{2}
	// next to check = 1 + the last number in the primes list
	// primeCounter := len(primes)

	// pPrime := primes[len(primes)-1] + 1

	for pPrime := primes[len(primes)-1] + 1; pPrime < lengthToSum; pPrime++ {

		for idx := 0; idx < len(primes); idx++ {
			// if divisible by a prime, it can't be a prime
			if pPrime%primes[idx] == 0 {
				goto newloop
			}
		}
		// pPrime is a Prime
		// pf("%d is a prime\n", pPrime)
		primes = append(primes, pPrime)
	newloop:
	}

	// sum the primes:
	sum := 0
	for i := 0; i < len(primes); i++ {
		sum += primes[i]
	}
	return sum
}

func sumOfPrimes2(lengthToSum int) int {
	// defer timeTrack(time.Now(), "sumOfPrimes2()") // Timer function

	primes := []int{2}
	// next to check = 1 + the last number in the primes list
	// primeCounter := len(primes)

	pPrime := primes[len(primes)-1] + 1
	// if pPrime is even
	if pPrime%2 == 0 {
		// make pprime odd
		pPrime += 1
	}

	for ; pPrime < lengthToSum; pPrime += 2 {

		for idx := 0; idx < len(primes); idx++ {
			// if divisible by a prime, it can't be a prime
			if pPrime%primes[idx] == 0 {
				goto newloop
			}
		}
		primes = append(primes, pPrime)
	newloop:
	}
	// sum the primes:
	sum := 0
	for i := 0; i < len(primes); i++ {
		sum += primes[i]
	}
	return sum
}

func sumOfPrimes3(lengthToSum int) int {
	// defer timeTrack(time.Now(), "sumOfPrimes3()") // Timer function

	primes := []int{2}
	// next to check = 1 + the last number in the primes list
	// primeCounter := len(primes)

	pPrime := primes[len(primes)-1] + 1
	// if pPrime is even
	if pPrime%2 == 0 {
		// make pprime odd
		pPrime += 1
	}

	for ; pPrime < lengthToSum; pPrime += 2 {
		// iterate through each prime value in the known primes list
		maxFactor := int(math.Sqrt(float64(pPrime))) // the larges possible factor for pPrime
		// idx = 1, as we never need to check the first prime (i.e. 2)
		for idx := 1; idx < len(primes); idx++ {
			// if remaining primes are bigger than maximum possible factor size
			if primes[idx] > maxFactor {
				// pPrime must be a prime, so can end loop
				break
			}
			if pPrime%primes[idx] == 0 {
				goto newloop
			}
		}
		// pPrime is a prime, so add to list of primes
		primes = append(primes, pPrime)
	newloop:
	}
	// sum the primes:
	sum := 0
	for i := 0; i < len(primes); i++ {
		sum += primes[i]
	}
	return sum
}

func main() {
	// p(sumOfPrimes1(10000)) // call the function being tested
	// p(sumOfPrimes2(10000)) // call the function being tested
	p(sumOfPrimes3(2000000)) // call the function being tested
}
