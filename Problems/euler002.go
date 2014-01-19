package main

import "fmt"
import "time"

// Used to time other functions to compare speed
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

func someName() {
	defer timeTrack(time.Now(), "Count1") // Timer function

	// Code body

}

func main() {
	someName() // call the function being tested
}
