// 215 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

// What is the sum of the digits of the number 21000?

package main

import (
	"fmt"
	"math"
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

func sumOfDigits1(power int) int {
	defer timeTrack(time.Now(), "sumOfDigits1()") // Timer function

	longNumber := []int{2, 0}
	num := 0
	carry := 0
	for i := 1; i < power; i++ {
		for j := 0; j < len(longNumber); j++ {
			num = longNumber[j]*2 + carry
			carry = 0
			longNumber[j] = num % 10
			// p("num: ", num)
			if num >= 10 {
				carry = num / 10
			}
		}
		if longNumber[len(longNumber)-1] != 0 {
			longNumber = append(longNumber, 0)
		}
	}

	// ADD ELEMENTS
	sum := 0
	for _, v := range longNumber {
		// p("v = ", v)
		sum += v
	}
	p("longNuber digits: ", len(longNumber)-1)
	return sum
}

// using a built in package function and the power of floats in Go
func sumOfDigits2(power float64) int {
	defer timeTrack(time.Now(), "sumOfDigits2()") // Timer function

	total := 0
	number := strconv.FormatFloat(math.Pow(2, power), 'f', 0, 64)

	for _, v := range number {
		digit, _ := strconv.Atoi(string(v))
		total += digit
	}
	return total
}

func main() {
	// p(sumOfDigits1(1000))
	p(sumOfDigits2(1000))

}
