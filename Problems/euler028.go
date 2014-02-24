/*
Starting with the number 1 and moving to the right in a clockwise direction a
5 by 5 spiral is formed as follows:

21 22 23 24 25
20  7  8  9 10
19  6  1  2 11
18  5  4  3 12
17 16 15 14 13

It can be verified that the sum of the numbers on the diagonals is 101.

What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral
formed in the same way?
*/

package main

import (
	"fmt"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// timeTrack is used for basic benchmarking in other functions
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

func diagonalSumOnSpiral(width int) int {
	defer timeTrack(time.Now(), "diagonalSumOnSpiral()")

	counter := 1
	sum := 1
	// run (1001+1)/2 loops
	for loop := 2; loop <= (width+1)/2; loop++ {
		// do 4 times for each loop
		inc := ((loop - 1) * 2)
		for i := 1; i <= 4; i++ {
			counter += inc
			sum += counter
		}
	}
	return sum
}

func main() {
	p(diagonalSumOnSpiral(1001))

}
