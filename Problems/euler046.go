/*
It was proposed by Christian Goldbach that every odd composite number can be
written as the sum of a prime and twice a square.

9 = 7 + 2×12
15 = 7 + 2×22
21 = 3 + 2×32
25 = 7 + 2×32
27 = 19 + 2×22
33 = 31 + 2×12

It turns out that the conjecture was false.

What is the smallest odd composite that cannot be written as the sum of a
prime and twice a square?
*/

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

func listPrimesAndOddComposites(max int) ([]int, []int) {
	// defer timeTrack(time.Now(), "mapOfPrimes") // Timer function

	sieve := make([]bool, max+2)
	primeList := make([]int, 0)
	oddComposites := make([]int, 0)

	// 2 is the first prime
	for i := 2; i <= max; i++ {
		if sieve[i] == false {
			// i should be a prime number, so add to list of primes
			primeList = append(primeList, i)
			for j := 2; i*j <= max; j++ {
				sieve[i*j] = true
			}
		} else if i%2 == 1 {
			oddComposites = append(oddComposites, i)
		}
	}
	return primeList, oddComposites
}

// findGoldbachException finds the first exception to Goldbach's other conjecture (if it exists) below the limit
func findGoldbachException() int {
	defer timeTrack(time.Now(), "findGoldbachException()")

	const LIMIT = 10000

	primeList, compositeList := listPrimesAndOddComposites(LIMIT)
	resultsMap := make(map[int]bool)

	// the largest number that the integer portion can be and still have a result less than 'lim'
	integerLimit := int(math.Sqrt(float64(LIMIT))) / 2
	// generate the resultMap where every
	for primeIdx := 0; primeIdx < len(primeList); primeIdx++ {
		for integer := 1; integer < integerLimit; integer++ {
			result := primeList[primeIdx] + 2*integer*integer
			// pf("%d + 2*%d^2 = %d\n", primeList[primeIdx], integer, result)
			// only bother with the odd results
			if result%2 != 0 {
				resultsMap[result] = true
			}
		}
	}

	// p(resultsMap)
	// p(compositeList)

	// compare the resultsMap and the compositeList, but ignore the first prime (2)
	for i := 0; i < len(compositeList); i++ {
		if resultsMap[compositeList[i]] == false {
			return compositeList[i]
		}
	}
	return 0

}

func main() {
	// p(listPrimesAndOddComposites(20))
	p(findGoldbachException())

}
