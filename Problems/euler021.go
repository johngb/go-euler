/*
Let d(n) be defined as the sum of proper divisors of n (numbers less than n
which divide evenly into n). If d(a) = b and d(b) = a, where a ≠ b, then a and
b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55
and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71
and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
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
	for i := 2; i <= maxDivisor; i++ {
		// if i is a factor of num
		if num%i == 0 {
			// add i and num/i to divisors
			divisors = append(divisors, i)
			divisors = append(divisors, num/i)
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

func isAmicable(num int) (bool, int) {

	num1DivisorSum := sumOfIntSlice(divisors(num))

	if sumOfIntSlice(divisors(num1DivisorSum)) == num {
		return true, num1DivisorSum
	}
	return false, 0
}

func sumOfAmicableNumbersUnderN(n int) int {
	defer timeTrack(time.Now(), "sumOfAmicableNumbersUnderN()")

	amicableMap := map[int]int{}
	sum, smaller, larger, secondNumber := 0, 0, 0, 0
	amicable := false
	for i := 1; i < n; i++ {
		amicable, secondNumber = isAmicable(i)
		if amicable {
			if i < secondNumber {
				smaller = i
				larger = secondNumber
			} else if i > secondNumber {
				smaller = secondNumber
				larger = i
				// they are equal, so move onto next number
			} else {
				continue
			}
			// if not already stored in the map
			if amicableMap[smaller] == 0 {
				amicableMap[smaller] = larger
				sum += larger + smaller
			}
		}
	}
	return sum
}

func main() {
	// p(divisors(9998))
	// p(isAmicable(284))
	p(sumOfAmicableNumbersUnderN(10000))

}
