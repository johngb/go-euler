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

// func functionName() {
// 	defer timeTrack(time.Now(), "functionName()")

// }

func main() {
	n := 3
	n = ^n
	p(n)

}
