/*
The primes 3, 7, 109, and 673, are quite remarkable. By taking any two primes
and concatenating them in any order the result will always be prime. For
example, taking 7 and 109, both 7109 and 1097 are prime. The sum of these four
primes, 792, represents the lowest sum for a set of four primes with this
property.

Find the lowest sum for a set of five primes for which any two primes
concatenate to produce another prime.
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

func primeSieve1(maxMap, maxList int) (map[int]bool, []int) {
	defer timeTrack(time.Now(), "primeSieve1()") // Timer function

	sieve := make([]bool, maxMap/2+1)
	// initialise the prime list with the only even prime
	primeList := []int{2}
	primeMap := map[int]bool{2: true}

	// 3 is the first odd prime
	for i := 3; i <= maxMap; i += 2 {
		if sieve[i/2] == false {
			// i should be a prime number, so add to list of primes
			// only add if less than maxList
			if i < maxList {
				primeList = append(primeList, i)
			}
			primeMap[i] = true
			// any odd * even = even, so avoid all even multiples
			for j := 3; i*j <= maxMap; j += 2 {
				sieve[i*j/2] = true
			}
		}
	}
	return primeMap, primeList
}

func findFifthPrime1(lim int) int {
	defer timeTrack(time.Now(), "findFifthPrime1()")

	// note that the limit of the list is not the same as the limit of the map
	pM, pL := primeSieve1(lim*lim, lim+200)

	pair := func(a, b int) bool {
		c, _ := strconv.Atoi(strconv.Itoa(pL[a]) + strconv.Itoa(pL[b]))
		if !pM[c] {
			return false
		}
		d, _ := strconv.Atoi(strconv.Itoa(pL[b]) + strconv.Itoa(pL[a]))
		if !pM[d] {
			return false
		}
		return true
	}
	for p1 := 1; pL[p1] < lim; p1++ {
		for p2 := p1 + 1; pL[p2] < lim; p2++ {
			if pair(p1, p2) {
				for p3 := p2 + 1; pL[p3] < lim; p3++ {
					if pair(p1, p3) && pair(p2, p3) {
						for p4 := p3 + 1; pL[p4] < lim; p4++ {
							if pair(p1, p4) && pair(p2, p4) && pair(p3, p4) {
								for p5 := p4 + 1; pL[p5] < lim; p5++ {
									if pair(p1, p5) && pair(p2, p5) && pair(p3, p5) && pair(p4, p5) {
										p("5set:", pL[p1], pL[p2], pL[p3], pL[p4], pL[p5])
										sum := pL[p1] + pL[p2] + pL[p3] + pL[p4] + pL[p5]
										return sum
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return -1
}

func main() {
	p(findFifthPrime1(10000))

}
