/*
The first ten terms in the sequence of convergents for e are:

2, 3, 8/3, 11/4, 19/7, 87/32, 106/39, 193/71, 1264/465, 1457/536, ...
The sum of digits in the numerator of the 10th convergent is 1+4+5+7=17.

Find the sum of digits in the numerator of the 100th convergent of the
continued fraction for e.
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

func sumDigitsInString(str string) int {

	counter := 0
	for i := 0; i < len(str); i++ {
		counter += int(str[i] - '0')
	}
	return counter
}

func doIt() int {
	defer timeTrack(time.Now(), "functionName()")

	const LIMIT = 100

	d := big.NewInt(1)
	n := big.NewInt(2)
	dOld := big.NewInt(1)
	bigZero := big.NewInt(0)
	bigA := big.NewInt(0)

	for i := 2; i <= LIMIT; i++ {
		dOld.Add(d, bigZero)
		if i%3 == 0 {
			bigA = big.NewInt(int64(2 * i / 3))
		} else {
			bigA = big.NewInt(1)
		}

		d.Add(n, bigZero)              // d = n
		n.Add(bigA.Mul(bigA, d), dOld) // n = bigA * d + dOld
	}
	p(n.String())
	return sumDigitsInString(n.String())
}

func main() {
	// p(sumDigitsInString("123"))
	p(doIt())

}
