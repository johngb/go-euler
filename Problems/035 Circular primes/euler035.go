/*
The number, 197, is called a circular prime because all rotations of the
digits: 197, 971, and 719, are themselves prime.

There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71,
73, 79, and 97.

How many circular primes are there below one million?
*/

package main

import (
	"fmt"
	"strconv"
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

func countCircularPrimes1(lim int) int {
	// defer timeTrack(time.Now(), "countCircularPrimes1()")

	// generate list of primes below 'lim' and above min
	primeMap := mapOfPrimes(2, lim)
	// create a count and set to 0
	count := 0
	// for each prime in the map
primeloop:
	for x, _ := range primeMap {
		k := x
		if k > lim {
			continue
		}
		revDigitList := make([]int, 0)
		// create a list of the digits of k (makes in reverse)
		for k/10 != 0 {
			revDigitList = append(revDigitList, k%10)
			k /= 10
		}
		revDigitList = append(revDigitList, k)
		// get the list in the right order
		digitList := make([]int, len(revDigitList))
		position := len(revDigitList) - 1
		for j := 0; j < len(revDigitList); j++ {
			digitList[position] = revDigitList[j]
			position--
		}

		// check each rotation to see if it's prime
		for i := 0; i < len(digitList); i++ {
			// create a rotated list
			digitList = append(digitList[1:len(digitList)], digitList[0])
			// convert list to a number
			pPrime := 0
			for j := 0; j < len(digitList); j++ {
				pPrime *= 10
				pPrime += digitList[j]
			}
			// check to see if it's a prime
			if primeMap[pPrime] == false {
				// goto outerloop
				continue primeloop
			}
		}
		count++
	}
	return count
}

// easier to understand, using strconv
func countCircularPrimes2(lim int) int {
	// defer timeTrack(time.Now(), "countCircularPrimes2()")

	// generate list of primes below 'lim' and above min
	primeMap := mapOfPrimes(2, lim)
	// create a count and set to 0
	count := 0
	// for each prime in the map
primeloop:
	for k, _ := range primeMap {

		kStr := strconv.Itoa(k)
		lenK := len(kStr)

		// check each rotation to see if it's prime
		for i := 0; i < lenK; i++ {
			// create a rotated list
			kStr = string(kStr[1:lenK]) + string(kStr[0])
			pPrime, _ := strconv.Atoi(kStr)
			// check to see if it's a prime
			if primeMap[pPrime] == false {
				// goto outerloop
				continue primeloop
			}
		}
		// increment the count
		count++
	}

	return count
}

func main() {

	p(countCircularPrimes1(1000000))
	p(countCircularPrimes2(1000000))

}
