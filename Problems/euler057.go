/*
It is possible to show that the square root of two can be expressed as an
infinite continued fraction.

√ 2 = 1 + 1/(2 + 1/(2 + 1/(2 + ... ))) = 1.414213...

By expanding this for the first four iterations, we get:

1 + 1/2 = 3/2 = 1.5
1 + 1/(2 + 1/2) = 7/5 = 1.4
1 + 1/(2 + 1/(2 + 1/2)) = 17/12 = 1.41666...
1 + 1/(2 + 1/(2 + 1/(2 + 1/2))) = 41/29 = 1.41379...

The next three expansions are 99/70, 239/169, and 577/408, but the eighth
expansion, 1393/985, is the first example where the number of digits in the
numerator exceeds the number of digits in the denominator.

In the first one-thousand expansions, how many fractions contain a numerator
with more digits than denominator?
*/

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

func bigIntPlusFraction(a int64, num *big.Int, den *big.Int) (*big.Int, *big.Int) {

	bigA := big.NewInt(a)
	newNum := new(big.Int)
	newNum.Mul(den, bigA)
	newNum.Add(newNum, num)
	return newNum, den
}

func bigIterations(iter int) (string, string) {
	num := big.NewInt(2)
	den := big.NewInt(1)
	for i := 1; i < iter; i++ {
		num, den = bigIntPlusFraction(int64(2), den, num) // note the switch between num and den here to invert the fraction
	}
	num, den = bigIntPlusFraction(int64(1), den, num)
	return num.String(), den.String()
}

// brute force approach without much simplification.
func bigEuler057(lim int) int {
	defer timeTrack(time.Now(), "bigEuler057()")

	count := 0

	for i := 1; i <= lim; i++ {
		strNum, strDen := bigIterations(i)
		// strNum := strconv.Itoa(num)
		// strDen := strconv.Itoa(den)

		if len(strNum) > len(strDen) {
			count++
		}
	}
	return count
}

// much faster implementation based on the fact that in each case
// if term n = x/y, term n+1 = 2*x+y / x+y
func bigEuler057a(lim int) int {
	defer timeTrack(time.Now(), "bigEuler057a()")
	counter := 0
	big2 := big.NewInt(2)

	num := big.NewInt(3)
	den := big.NewInt(2)

	temp := big.NewInt(0)

	for i := 1; i < lim; i++ {
		if num.String() < den.String() {
			counter++
		}
		temp.Mul(den, big2)
		den.Add(num, den)
		num.Add(num, temp)
	}
	return counter
}

func main() {
	p(bigEuler057(1000))
	p(bigEuler057a(1000))
}
