/*

The arithmetic sequence, 1487, 4817, 8147, in which each of the terms
increases by 3330, is unusual in two ways:

(i) each of the three terms areprime, and,
(ii) each of the 4-digit numbers are permutations of one another.

There are no arithmetic sequences made up of three 1-, 2-, or 3-digit primes,
exhibiting this property, but there is one other 4-digit increasing sequence.

What 12-digit number do you form by concatenating the three terms in this
sequence? */

package main

import (
	"fmt"
	"strconv"
	"strings"
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
	defer timeTrack(time.Now(), "MapOfPrimes()") // Timer function

	sieve := make([]bool, max+1)
	// initialise the prime list with the only even prime
	primeList := []int{2}

	// 3 is the first odd prime
	for i := 3; i <= max; i += 2 {
		if sieve[i] == false {
			// i should be a prime number, so add to list of primes
			primeList = append(primeList, i)
			for j := 2; i*j <= max; j++ {
				sieve[i*j] = true
			}
		}
	}
	primeMap := make(map[int]bool)
	for i := 0; i < len(primeList); i++ {
		if primeList[i] >= min {
			primeMap[primeList[i]] = true
		}
	}
	return primeMap
}

func isPermutation(a, b int) bool {

	strA := strconv.Itoa(a)
	strB := strconv.Itoa(b)

	for i := 0; i < len(strA); {
		char := string(strA[0])
		if strings.Contains(strB, char) {
			//remove that character once from both strings
			strA = strings.Replace(strA, char, "", 1)
			strB = strings.Replace(strB, char, "", 1)
		} else {
			return false
		}
	}
	return true
}

func findPrimeSeq() int {
	defer timeTrack(time.Now(), "findPrimeSeq()")

	primeMap := mapOfPrimes(1000, 10000)
	const gap = 3330

	for i := 1000; i < 10000; i++ {
		// if not a prime, get next number
		if primeMap[i] == false {
			continue
		}
		if primeMap[i+gap] == false {
			continue
		}
		if primeMap[i+2*gap] == false {
			continue
		}
		// at this point we should have three numbers separated by gap that
		// are all prime, so we need to check if they are permutations of one
		// another
		if isPermutation(i, i+gap) && isPermutation(i, i+gap*2) && i != 1487 {
			return i*1e8 + (i+gap)*1e4 + (i + 2*gap)
		}
	}
	return 0
}

func main() {

	// p(isPermutation(1323, 1233))
	p(findPrimeSeq())

}
