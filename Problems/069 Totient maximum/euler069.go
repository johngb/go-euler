/*
Euler's Totient function, φ(n) [sometimes called the phi function], is used to
determine the number of numbers less than n which are relatively prime to n.
For example, as 1, 2, 4, 5, 7, and 8, are all less than nine and relatively
prime to nine, φ(9)=6.

n	Relatively Prime	φ(n)	n/φ(n)
2	1					1		2
3	1,2					2		1.5
4	1,3					2		2
5	1,2,3,4				4		1.25
6	1,5					2		3
7	1,2,3,4,5,6			6		1.1666...
8	1,3,5,7				4		2
9	1,2,4,5,7,8			6		1.5
10	1,3,7,9				4		2.5
It can be seen that n=6 produces a maximum n/φ(n) for n ≤ 10.

Find the value of n ≤ 1,000,000 for which n/φ(n) is a maximum.
*/

// JGB: Basically it amounts to finding the number < limit with the most prime
// factors.

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

func primeSieve(max int) []int {
	// defer timeTrack(time.Now(), "primeSieve") // Timer function

	sieve := make([]bool, max/2+1)
	// initialise the prime list with the only even prime
	primeList := []int{2}

	// 3 is the first odd prime
	for i := 3; i <= max; i += 2 {
		if sieve[i/2] == false {
			// i should be a prime number, so add to list of primes
			primeList = append(primeList, i)
			// any odd * even = even, so avoid all even multiples
			for j := 3; i*j <= max; j += 2 {
				sieve[i*j/2] = true
			}
		}
	}
	return primeList
}

func euler069() int {
	defer timeTrack(time.Now(), "euler069()")

	result := 1
	primes := primeSieve(100)
	const limit = 1000000

	for i := 0; result*primes[i] <= limit; i++ {
		result *= primes[i]
	}
	return result
}

func main() {
	p(euler069())

}
