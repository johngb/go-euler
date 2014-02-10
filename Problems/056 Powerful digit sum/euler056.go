// A googol (10^100) is a massive number: one followed by one-hundred zeros;
// 100^100 is almost unimaginably large: one followed by two-hundred zeros.
// Despite their size, the sum of the digits in each number is only 1.

// Considering natural numbers of the form, a^b, where a, b < 100, what is the
// maximum digital sum?

package main

import (
	"fmt"
	"math/big"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// timeTrack is used for basic benchmarking in other functions
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

func digitSumOfPower(a, b int) int {
	// defer timeTrack(time.Now(), "digitSumOfPower()")

	bigA := big.NewInt(int64(a))
	bigB := big.NewInt(int64(b))
	bigAns := big.NewInt(0)
	bigAns.Exp(bigA, bigB, nil)

	digitStr := bigAns.String()

	counter := 0
	for i := 0; i < len(digitStr); i++ {
		counter += int(digitStr[i] - '0')
	}
	return counter
}

func maxDigitSumOfPower(lim int) int {
	defer timeTrack(time.Now(), "maxDigitSumOfPower()")

	maxDigitSum := 0
	// only need to consider the top 90% of values for largest digits.  This
	// is faster, but if you want to be safer, start with 1.  Both achieve the
	// same result
	startingPoint := int(lim * 9 / 10)

	for a := startingPoint; a < lim; a++ {
		for b := startingPoint; b < lim; b++ {
			digitSum := digitSumOfPower(a, b)
			if digitSum > maxDigitSum {
				maxDigitSum = digitSum
			}
		}
	}
	return maxDigitSum
}

func main() {

	p(maxDigitSumOfPower(100))

}
