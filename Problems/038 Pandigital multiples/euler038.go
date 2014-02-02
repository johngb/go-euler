/*
Take the number 192 and multiply it by each of 1, 2, and 3:

192 × 1 = 192 192 × 2 = 384 192 × 3 = 576 By concatenating each product we get
the 1 to 9 pandigital, 192384576. We will call 192384576 the concatenated
product of 192 and (1,2,3)

The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4, and
5, giving the pandigital, 918273645, which is the concatenated product of 9
and (1,2,3,4,5).

What is the largest 1 to 9 pandigital 9-digit number that can be formed as the
concatenated product of an integer with (1,2, ... , n) where n > 1?
*/

package main

import (
	"fmt"
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

// isPandigital returns true if the string 'str' is a pandigital number
func isPandigital(str string) bool {
	// defer timeTrack(time.Now(), "isPandigital()")

	if len(str) != 9 {
		return false
	}

	for i := 1; i <= 9; i++ {
		// if the string doesn't contain one of the numbers
		if !strings.Contains(str, strconv.Itoa(i)) {
			return false
		}
	}
	return true
}

func largestPandigital() int {
	defer timeTrack(time.Now(), "main()")

	largest := 0
	tempstring := ""
	tempint := 0

	// number that will be multiplied j times
	for i := 1; i < 10000; i++ {
		str := ""

		for j := 2; j <= 9 && len(str) < 9; j++ {
			for n := 1; n <= j; n++ {
				tempstring = strconv.Itoa(n * i)
				str += tempstring

			}
			if isPandigital(str) {
				tempint, _ = strconv.Atoi(str)
				if tempint > largest {
					largest = tempint
				}
			}
		}
	}
	return largest
}

func main() {

	p(isPandigital("123456789"))
	p(largestPandigital())

}
