/*
The fraction 49/98 is a curious fraction, as an inexperienced mathematician in
attempting to simplify it may incorrectly believe that 49/98 = 4/8, which is
correct, is obtained by cancelling the 9s.

We shall consider fractions like, 30/50 = 3/5, to be trivial examples.

There are exactly four non-trivial examples of this type of fraction, less
than one in value, and containing two digits in the numerator and denominator.

If the product of these four fractions is given in its lowest common terms,
find the value of the denominator.
*/

//JGB: Check out http://www.roberthorvick.com/2012/07/20/project-euler-problem-33-explained/ for an explanation

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

func prodOfDigitCancellingFractionDenominator() int {
	defer timeTrack(time.Now(), "sumOfDigitCancellingFractions()")

	var i, j, k, ik, kj int
	result := 1.0

	// k is the number that we will cancel out, while i and j are the second
	// digits in the numerator and denominator.  We exclude the values of 0 as
	// they are trivial
	for k = 1; k <= 9; k++ {
		for i = 1; i <= 9; i++ {
			for j = 1; j <= 9; j++ {
				ik = i*10 + k
				kj = k*10 + j

				if ik < kj {

					// if ik/jk = i/j, then j*ik = i*kj (by maths expansion)
					if j*ik == i*kj {
						p(ik, "/", kj)
						temp := float64(kj) / float64(ik)
						result *= temp
					}
				}

			}
		}

	}
	return int(result)
}

// func isValid(num, den int) bool {
// 	defer timeTrack(time.Now(), "isValid()")

func main() {
	p(prodOfDigitCancellingFractionDenominator())

}
