/*
The 5-digit number, 16807=7^5, is also a fifth power. Similarly, the 9-digit
number, 134217728=8^9, is a ninth power.

How many n-digit positive integers exist which are also an nth power?
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

// timeTrack is used for basic benchmarking in other functions
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

func length(num int) int {
	str := strconv.Itoa(num)
	return len(str)
}

func allZeroes(slice []int) bool {

	for i := 0; i < len(slice); i++ {
		if slice[i] != 0 {
			return false
		}
	}
	return true
}

func doIt() int {
	defer timeTrack(time.Now(), "doIt()")

	power := 2
	previouscount := 9
	count := 9 // the first 9 digits to the power of 1 are 1 digit long

	// want the first element to be 1, so fill the 0th element
	arr := []*big.Int{}
	for i := int64(0); i < 10; i++ {
		bigNum := big.NewInt(i)
		arr = append(arr, bigNum)
	}

	// make initial slice
	for i := 1; ; i++ {
		for j := 1; j < 10; j++ {
			arr[j].Mul(arr[j], big.NewInt(int64(j)))
			if len(arr[j].String()) == power {
				count++
				// if it's smaller, then it will keep getting smaller, and if
				// bigger, it will keep getting bigger, so basically remove it.
				// This also provides a way to check if it's done.
			} else {
				arr[j] = big.NewInt(0)
			}
		}
		if previouscount == count {
			return count
		}
		previouscount = count
		power++
	}
	return count
}

func main() {

	p(doIt())

}
