/*
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/
package main

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// Used to time other functions to compare speed
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

func factorial(n int64) *big.Int {
	if n < 0 {
		return big.NewInt(1)
	}
	if n == 0 {
		return big.NewInt(1)
	}
	bigN := big.NewInt(n)
	return bigN.Mul(bigN, factorial(n-1))
}

// convert a char to an int.  No checking int place
func ord(c byte) int {
	return int(c - '0')
}

func sumOfFactorialDigits1(n int) int {
	defer timeTrack(time.Now(), "sumOfFactorialDigits1()")

	digit, result := 0, 0
	st := factorial(int64(n)).String()
	for i := range st {
		digit = ord(st[i])
		result += digit
	}
	return result
}

func sumOfFactorialDigits2(lim int) int {
	defer timeTrack(time.Now(), "sumOfFactorialDigits2()")

	mul := big.NewInt(1)
	for i := 2; i <= lim; i++ {
		mul.Mul(mul, big.NewInt(int64(i)))
	}
	result, digit := 0, 0
	for i := range mul.String() {
		digit, _ = strconv.Atoi(string(mul.String()[i]))
		result += digit
	}
	return result
}

func main() {

	p(sumOfFactorialDigits1(100))
	p(sumOfFactorialDigits2(100))

}
