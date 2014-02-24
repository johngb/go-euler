/*
The series, 11 + 22 + 33 + ... + 1010 = 10405071317.

Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.
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

func selfPower() string {
	defer timeTrack(time.Now(), "selfPower()")

	sum := big.NewInt(0)
	power := big.NewInt(1)
	number := big.NewInt(0)

	for i := 1; i <= 1000; i++ {
		number = big.NewInt(int64(i))
		power.Exp(number, number, nil)
		sum.Add(sum, power)
	}

	str := sum.String()
	return str[len(str)-10 : len(str)]

}

func main() {

	p(selfPower())

}
