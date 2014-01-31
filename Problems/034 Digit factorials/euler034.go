/*
145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of
their digits.

Note: as 1! = 1 and 2! = 2 are not sums they are not included.
*/

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

func factorial(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 1
	}
	fact := 1
	for i := 2; i <= n; i++ {
		fact *= i
	}
	return fact
}

func findFactorialDigitSumNumbers() int {
	defer timeTrack(time.Now(), "findFactorialDigitSumNumbers()")

	results := []int{}

	factorials := make([]int, 10)
	factorials[0] = 1
	for a := 1; a <= 9; a++ {
		factorials[a] = a * factorials[a-1]
	}

	for i := 10; i < factorials[9]*10; i++ {
		sum := 0
		j := i
		for j > 9 {
			sum += factorials[j%10]
			j /= 10
		}
		sum += factorials[j]

		if sum == i {
			results = append(results, i)
		}
	}

	// add results
	finalSum := 0
	for _, v := range results {
		finalSum += v
	}

	return finalSum
}

func main() {

	p(findFactorialDigitSumNumbers())

}
