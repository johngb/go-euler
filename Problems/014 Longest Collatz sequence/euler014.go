/*The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.*/

package main

import (
	"fmt"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// Used to time other functions to compare speed
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

// collatz1 returns the collatz sequence starting with 'n'
func collatzSequence1(n int) []int {
	// defer timeTrack(time.Now(), "collatzSequence1") // Timer function

	sequence := []int{n}
	for {
		// if n is even
		if n%2 == 0 {
			n = n / 2
			// otherwise n is odd
		} else {
			n = 3*n + 1
		}
		sequence = append(sequence, n)
		if n == 1 {
			return sequence
		}
	}
}

// collatz1 returns the collatz sequence starting with 'n'
func collatzLenght1(n int) int {
	// defer timeTrack(time.Now(), "collatzLenght1") // Timer function

	sequence := 1
	for {
		// if n is even
		if n%2 == 0 {
			n = n / 2
			// otherwise n is odd
		} else {
			n = 3*n + 1
		}
		sequence++
		if n == 1 {
			return sequence
		}
	}
}

// findCollatzSequence1 finds the starting number which produces
// a Collatz sequence 'length' long
// this is a brute force approach, and very inefficient
func findCollatzSequence1(length int) int {
	defer timeTrack(time.Now(), "findCollatzSequence1()") // Timer function

	for i := 2; ; i++ {
		if len(collatzSequence1(i)) == length {
			return i
		}
	}
	return 0 // error if this happens
}

// findCollatzSequence2 finds the starting number which produces
// a Collatz sequence >= 'length' long
// this is a brute force approach, and very inefficient
func findCollatzSequence2(length int) int {
	defer timeTrack(time.Now(), "findCollatzSequence2()") // Timer function

	// map to hold the sequence counts for known sequences
	sequenceCounts := make(map[int]int)
	// seed the array
	sequenceCounts[0] = 1
	n := 0 // itermediate number
	iterations := 0
	sequenceLength := 0
	for i := 2; ; i++ {
		n = i
		iterations, sequenceLength = 0, 0
		// pf("i: %d, n: %d, iterations: %d, sequenceLength: %d\n", i, n, iterations, sequenceLength)
		for {
			if n%2 == 0 {
				n = n / 2
			} else {
				n = 3*n + 1
			}
			iterations++
			// CHECKS

			// if there is a sequence already stored, we know the length of i
			if sequenceCounts[n-1] != 0 {
				sequenceLength = iterations + sequenceCounts[n-1]

				// store the length of i for future sequences
				sequenceCounts[i-1] = sequenceLength

				// check for end condition
				if sequenceLength >= length {
					return i
				}
				break
			}
		}
	}
	return 0 // error if this happens
}

// collatzWithMaxLength1 returns the number below 'n' with the longest collatz sequence
func collatzWithMaxLength1(n int) int {
	defer timeTrack(time.Now(), "collatzWithMaxLength1") // Timer function

	longestSequence := 0
	numberWithlongestSequence := 0
	sequence := 0
	for i := 1; i < n; i++ {
		sequence = collatzLenght1(i)
		if sequence > longestSequence {
			longestSequence = sequence
			numberWithlongestSequence = i
		}
	}
	// p(longestSequence)
	return numberWithlongestSequence
}

// collatzWithMaxLength2 returns the number below 'num' with the longest collatz sequence
// tried to use a map to reduce the number of times that a repeat sequence is checked, but was slower
func collatzWithMaxLength2(num int) int {
	defer timeTrack(time.Now(), "collatzWithMaxLength2()") // Timer function

	// map to hold the sequence counts for known sequences
	sequenceCounts := make(map[int]int)
	// seed the map
	sequenceCounts[0] = 1
	n := 0 // itermediate number
	iterations := 0
	sequenceLength := 0

	longestSequence := 0
	numberWithlongestSequence := 0

	for i := 2; i < num; i++ {
		n = i
		iterations, sequenceLength = 0, 0
		for {
			if n%2 == 0 {
				n = n / 2
			} else {
				n = 3*n + 1
			}
			iterations++

			// CHECKS
			// if there is a sequence already stored, we know the length of i
			if sequenceCounts[n-1] != 0 {
				sequenceLength = iterations + sequenceCounts[n-1]
				// pf("i: %d, n: %d, iterations: %d, sequenceLength: %d\n", i, n, iterations, sequenceLength)

				// store the length of i for future sequences
				sequenceCounts[i-1] = sequenceLength
				if sequenceLength > longestSequence {
					longestSequence = sequenceLength
					numberWithlongestSequence = i
				}
				break
			}
		}
	}
	return numberWithlongestSequence
}

func main() {

	const n = 100000
	// ans := findCollatzSequence2(n)
	// p(ans)
	// p("length collatz: ", len(collatz1(ans)))
	// p(collatz1(ans))
	// p(collatzLenght1(n))
	p(collatzWithMaxLength1(n))
	p(collatzWithMaxLength2(n))

}
