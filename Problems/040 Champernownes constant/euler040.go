/*
An irrational decimal fraction is created by concatenating the positive
integers:

0.123456789101112131415161718192021...

It can be seen that the 12th digit of the fractional part is 1.

If dn represents the nth digit of the fractional part, find the value of the
following expression.

d1 × d10 × d100 × d1000 × d10000 × d100000 × d1000000
*/

package main

import (
	"fmt"
	"strconv"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// timeTrack is used for basic benchmarking in other functions
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

func ord(c byte) int {
	return int(c - '0')
}

// Note: This is a brute force method that has potential for lots of
// optimisation.  I'm leaving that for now as it is a fairly straightforward
// exercise
func findChamConst() int {
	defer timeTrack(time.Now(), "findChamConst()")

	longStr := ""
	str := ""
	for i := 1; len(longStr) <= 1000000; i++ {
		str = strconv.Itoa(i)
		longStr += str
	}
	return ord(longStr[1-1]) * ord(longStr[10-1]) * ord(longStr[100-1]) * ord(longStr[1000-1]) * ord(longStr[10000-1]) * ord(longStr[100000-1]) * ord(longStr[1000000-1])
}

func main() {

	p(findChamConst())

}
