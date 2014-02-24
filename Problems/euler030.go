// Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.

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

// convert a char to an int.  No checking int place
func ord(c byte) int {
	return int(c - '0')
}

func findSumOfFifthPowerDigits() int {
	defer timeTrack(time.Now(), "findSumOfFifthPowerDigits()")

	fifthPowerNumbers := []int{}

	digits := ""
	for i := 2; i < 1000000; i++ {
		sum := 0
		digits = strconv.Itoa(i)
		for j := 0; j < len(digits); j++ {
			t := ord(digits[j])
			sum += t * t * t * t * t
		}
		if sum == i {

			fifthPowerNumbers = append(fifthPowerNumbers, i)
		}
	}
	sumNumbers := 0
	for _, n := range fifthPowerNumbers {
		sumNumbers += n
	}

	return sumNumbers
}

func main() {
	p(findSumOfFifthPowerDigits())
}
