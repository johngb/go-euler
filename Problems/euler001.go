package main

import (
	"fmt"
	"time"
)

// Used to time other functions to compare speed
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

// This function simply counts from 1 to 999 and adds any numbers that are
// factors of the current number.  It's basically a brute force approach
func Count1(countTo int) {
	// Timer function
	defer timeTrack(time.Now(), "Count1")

	sum := 0
	for i := 1; i < countTo; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Printf("Count1: Sum = %d\n", sum)
	return
}

// This function uses the principle that Sum(3||5) = Sum(3) + Sum(5) - Sum(3*5)
// It's much longer than the first solution, but is also about 10 times faster
func Count2(countTo int) {
	// Timer function
	defer timeTrack(time.Now(), "Count2")

	// Sum(3)
	sum := 0
	for i := 3; i <= int(countTo/3)*3; i += 3 {
		sum += i
	}

	// Sum(5)
	for i := 5; i <= int(countTo/5)*5; i += 5 {
		sum += i
	}

	// Sum(15)
	for i := 15; i <= int(countTo/15)*15; i += 15 {
		sum -= i
	}

	fmt.Printf("Count2: Sum = %d\n", sum)
	return
}

func main() {
	Count1(1000)
	Count2(1000)
}
