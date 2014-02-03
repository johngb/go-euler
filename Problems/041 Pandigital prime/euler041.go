/*
We shall say that an n-digit number is pandigital if it makes use of all the
digits 1 to n exactly once. For example, 2143 is a 4-digit pandigital and is
also prime.

What is the largest n-digit pandigital prime that exists?
*/

package main

import (
	"fmt"
	"math"
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

func lenBase10(a int) int {

	count := 0
	for a > 0 {
		count += 1
		a /= 10
	}
	return count
}

// Note: Using a map and trying to handle it without string functions
func isPandigitalN1(a int) bool {
	// defer timeTrack(time.Now(), "isPandigitalN1()")

	length := 0
	digitMap := make(map[int]bool)
	for a > 0 {
		digit := a % 10
		if digitMap[digit] {
			return false
		} else {
			digitMap[digit] = true
		}
		length += 1
		a /= 10
	}

	for i := 1; i <= length; i++ {
		// if the string doesn't contain one of the numbers
		if !digitMap[i] {
			return false
		}
	}
	return true
}

// Note: Using string functions.  Turns out to be faster than isPandigitalN1.
func isPandigitalN2(a int) bool {
	// defer timeTrack(time.Now(), "isPandigital()")

	str := strconv.Itoa(a)
	length := len(str)

	for i := 1; i <= length; i++ {
		// if the string doesn't contain one of the numbers
		if !strings.Contains(str, strconv.Itoa(i)) {
			return false
		}
	}
	return true
}

// Note: Using an array and no string functions.  Much faster with the array
func isPandigitalN3(a int) bool {
	// defer timeTrack(time.Now(), "isPandigitalN1()")

	length := 0
	var digitMap [10]bool
	for a > 0 {
		digit := a % 10
		if digitMap[digit] {
			return false
		} else {
			digitMap[digit] = true
		}
		length += 1
		a /= 10
	}

	for i := 1; i <= length; i++ {
		// if the string doesn't contain one of the numbers
		if !digitMap[i] {
			return false
		}
	}
	return true
}

func largestPandigitalPrime1() int {
	defer timeTrack(time.Now(), "largestPandigitalPrime1()")

	for i := 9876543; i > 0; i -= 2 {
		if isPandigitalN3(i) && isPrime(i) {
			return i
		}
	}
	return -1
}

func largestPandigitalPrime2() int {
	defer timeTrack(time.Now(), "largestPandigitalPrime2()")

	for i := 9876543; i > 0; i -= 2 {
		if isPandigitalN2(i) && isPrime(i) {
			return i
		}
	}
	return -1
}

func main() {
	p(largestPandigitalPrime1())
	p(largestPandigitalPrime2())
	// p(isPandigitalN(1234576789))

}
