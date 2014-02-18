/*
All square roots are periodic when written as continued fractions and can be written in the form:

** Left out here as it wasn't useful.  Rather see:
http://en.wikipedia.org/wiki/Methods_of_computing_square_roots#Continued_fraction_expansion and
http://web.math.princeton.edu/mathlab/jr02fall/Periodicity/alexajp.pdf

How many continued fractions for N ≤ 10000 have an odd period?
*/

// JGB: http://www.mathblog.dk/project-euler-continued-fractions-odd-period/ was useful in understanding the problem.

package main

import (
	"fmt"
	"math"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// timeTrack is used for basic benchmarking in other functions
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

func findOddPeriodContinuedFractions(lim int) int {
	defer timeTrack(time.Now(), "findOddPeriodContinuedFractions()")

	count := 0
	for s := 2; s <= lim; s++ {
		a0 := int(math.Sqrt(float64(s)))
		// if n is a perfect square, continue with the next number
		if a0*a0 == s {
			continue
		}

		period := 0
		d := 1
		m := 0
		a := a0

		for {
			// for the algorithm see http://en.wikipedia.org/wiki/Methods_of_computing_square_roots#Example.2C_square_root_of_114_as_a_continued_fraction
			m = d*a - m
			d = (s - m*m) / d
			a = (a0 + m) / d
			period++
			// for the end condition see corrolary 3.3 in http://web.math.princeton.edu/mathlab/jr02fall/Periodicity/alexajp.pdf
			if a == 2*a0 {
				break
			}
		}
		if period%2 == 1 {
			count++
		}
	}
	return count
}

func findOddPeriodContinuedFractions1(lim int) int {
	defer timeTrack(time.Now(), "findOddPeriodContinuedFractions1()")

	count := 0
	for s := 2; s <= lim; s++ {
		a0 := int(math.Sqrt(float64(s)))
		// if n is a perfect square, continue with the next number
		if a0*a0 == s {
			continue
		}

		period := 0
		d := 1
		m := 0
		a := a0

		// for the end condition see corrolary 3.3 in http://web.math.princeton.edu/mathlab/jr02fall/Periodicity/alexajp.pdf
		for ; a != 2*a0; period++ {
			// for the algorithm see http://en.wikipedia.org/wiki/Methods_of_computing_square_roots#Example.2C_square_root_of_114_as_a_continued_fraction
			m = d*a - m
			d = (s - m*m) / d
			a = (a0 + m) / d
		}
		// if the period is odd, increment count
		if period%2 == 1 {
			count++
		}
	}
	return count
}

func main() {
	p(findOddPeriodContinuedFractions(10000))
	p(findOddPeriodContinuedFractions1(10000))
}
