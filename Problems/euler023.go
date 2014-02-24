/*
A perfect number is a number for which the sum of its proper divisors is
exactly equal to the number. For example, the sum of the proper divisors of 28
would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than
n and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest
number that can be written as the sum of two abundant numbers is 24. By
mathematical analysis, it can be shown that all integers greater than 28123
can be written as the sum of two abundant numbers. However, this upper limit
cannot be reduced any further by analysis even though it is known that the
greatest number that cannot be expressed as the sum of two abundant numbers is
less than this limit.

Find the sum of all the positive integers which cannot be written as the sum
of two abundant numbers.
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

func divisors(num int) []int {
	// defer timeTrack(time.Now(), "divisors()")

	divisors := []int{1}
	maxDivisor := int(math.Sqrt(float64(num)))
	// p("maxDivisor = ", maxDivisor)
	for i := 2; i <= maxDivisor; i++ {
		// if i is a factor of num
		if num%i == 0 {
			// add i to divisors
			divisors = append(divisors, i)
			// if i and num/i arent' equal, then add both
			if i != num/i {
				divisors = append(divisors, num/i)
			}
		}
	}
	return divisors
}

func sumOfIntSlice(s []int) int {
	count := 0
	for i := range s {
		count += s[i]
	}
	return count
}

// isAbundant determines whether a positive integer is an abundant number
func isAbundant(num int) bool {
	// defer timeTrack(time.Now(), "isAbundant()")

	if sumOfIntSlice(divisors(num)) > num {
		return true
	}
	return false
}

func makeMapOfAbundantNumbers(lim int) map[int]bool {
	// defer timeTrack(time.Now(), "makeMapOfAbundantNumbers()")

	numberMap := make(map[int]bool)

	for i := 1; i < lim; i++ {
		if isAbundant(i) {
			numberMap[i] = true
		}
	}
	return numberMap
}

func sumNumNotComposed2AbundantNumbers() int {
	defer timeTrack(time.Now(), "sumNumNotComposed2AbundantNumbers()")

	const MAX_NUMBER = 28123

	notComposedList := make([]int, 0)
	abundantMap := makeMapOfAbundantNumbers(MAX_NUMBER)

	// iterate over integers less than MAX_NUMBER
	for i := 1; i < MAX_NUMBER; i++ {

		// check to see if it is divisible by any abundant number
		for j := range abundantMap {
			// only check if i > j to reduce checks
			if i > j {
				if abundantMap[i-j] == true {
					// 'i' can be expressed as the sum of two abundant numbers, so next number
					goto newloop
				}
			}
		}
		notComposedList = append(notComposedList, i)
	newloop:
	}
	return sumOfIntSlice(notComposedList)
}

func main() {

	p(sumNumNotComposed2AbundantNumbers())
}
