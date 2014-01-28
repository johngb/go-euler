/*
The Fibonacci sequence is defined by the recurrence relation:

Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
Hence the first 12 terms will be:

F1 = 1
F2 = 1
F3 = 2
F4 = 3
F5 = 5
F6 = 8
F7 = 13
F8 = 21
F9 = 34
F10 = 55
F11 = 89
F12 = 144
The 12th term, F12, is the first term to contain three digits.

What is the first term in the Fibonacci sequence to contain 1000 digits?
*/
package main

import (
	"fmt"
	"math/big"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// Used to time other functions to compare speed
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

func findFactorialGreaterThanNumDigits(num int) int {
	defer timeTrack(time.Now(), "findFactorialGreaterThanNumDigits()") // Timer function

	a := big.NewInt(1)
	b := big.NewInt(1)
	temp := big.NewInt(1)
	iteration := 2 // corresponding to the value of 'b'

	// steps to follow for iterations
	for len(b.String()) < num {
		temp.Set(a) // equivalent to temp = a
		a.Set(b)
		b.Add(temp, b)
		iteration++
	}
	return iteration
}

func main() {
	p(findFactorialGreaterThanNumDigits(1000))
}
