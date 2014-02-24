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

// primeList returns a list of primes <= num, and their product
func makePrimeList(num int) ([]int, int) {
	// defer timeTrack(time.Now(), "primeList()") // Timer function

	// initialise an array of with num elements
	sieve := make([]bool, num+2)
	// intialise slice to contain prime numbers
	primes := make([]int, 0)
	// counter to track the product of the primes
	prod := 1

	// index = 2
	for i := 2; i <= num; i++ {

		if sieve[i] == false {

			// i should be a prime number, so add to list of primes
			primes = append(primes, i)
			// add primes to product of primes
			prod *= i

		}

		for j := 2; i*j <= num; j++ {
			sieve[i*j] = true
		}
	}

	return primes, prod
}

// primeMask creates a mask of non-primes of the size passed in 'num',
// where 'false' indicates that a number at that index is not a prime.
// in general this method will not apply for the primes below 'lim', but a list
// of them will be given in 'primeList'
func primeMask(primeList []int, lim int) (_sieve []bool, _masklimit int, _primeList []int) {
	// defer timeTrack(time.Now(), "primeMask()") // Timer function

	// initialise an array of with lim elements
	sieve := make([]bool, lim)
	// sieve[0] = true
	// iterate through the elements of the primeList
	for i := 0; i < len(primeList); i++ {
		// set all multiples of primeList[i] = 'true'
		for j := 1; primeList[i]*j <= lim; j++ {
			// pf("primeList[%d] = %d, j = %d\n", i, primeList[i], j)
			sieve[primeList[i]*j-1] = true

		}
	}
	return sieve, lim, primeList
}

func sumOfPrimes3(lengthToSum int) int {
	defer timeTrack(time.Now(), "sumOfPrimes3()") // Timer function

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

// significantly more complicated method, and only marginally faster for large prime lists.  Recommend sumOfPromes3 instead.
func sumOfPrimes4(limit int, tuning int) int {
	defer timeTrack(time.Now(), "sumOfPrimes4()") // Timer function

	mask, maskSize, maskedPrimeList := primeMask(makePrimeList(tuning))

	// start off the list with the list of primes masked in the mask
	primeList := maskedPrimeList

	pPrime := maskedPrimeList[len(maskedPrimeList)-1] + 1
	// if pPrime is even
	if pPrime%2 == 0 {
		// make pprime odd
		pPrime += 1
	}

	for ; pPrime < limit; pPrime += 2 {

		// first check the mask to see if it's NOT prime
		// pf("pPrime = %d, maskSize = %d, pPrime%%maskSize = %d\n", pPrime, maskSize, pPrime%maskSize)
		// pf("Mask[%d] = %t\n", pPrime%maskSize, mask[pPrime%maskSize-1])

		// if the mask value == true, then it isn't a prime
		if mask[pPrime%maskSize-1] == true {
			// p("--- skipping some checks ---")
			continue
		}
		// otherwise, it may be a prime, so need to check

		// iterate through each prime value in the known primeList list
		maxFactor := int(math.Sqrt(float64(pPrime))) // the larges possible factor for pPrime
		// idx = 1, as we never need to check the first prime (i.e. 2)
		for idx := 1; idx < len(primeList); idx++ {
			// if remaining primeList are bigger than maximum possible factor size
			if primeList[idx] > maxFactor {
				// pPrime must be a prime, so can end loop
				break
			}
			// if one of the primes in primeList is a factor, then pPrimes is not a prime
			if pPrime%primeList[idx] == 0 {
				goto newloop
			}
		}
		// pPrime is a prime, so add to list of primeList
		primeList = append(primeList, pPrime)
	newloop:
	}
	// sum the primeList:
	sum := 0
	for i := 0; i < len(primeList); i++ {
		sum += primeList[i]
	}
	return sum
}

func main() {

	// t1, t2, t3 := primeMask(makePrimeList(6))
	// p(t1)
	// p(t2)
	// p(t3)
	// p("----------------")
	const n = 2000000
	p(sumOfPrimes3(n))
	p(sumOfPrimes4(n, 10))
	p(sumOfPrimes5(n))
	// // p("-------------")
}
