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

func functionName() {
	defer timeTrack(time.Now(), "functionName()")

}

// pure brute force.  Could easily be optimised
func numberOfWaysToMake2Pounds() int {
	// defer timeTrack(time.Now(), "numberOfWaysToMake2Pounds1()")

	sum := 0
	// for 2 pounds + nothing else.
	vaildCombiCount := 1

	for f1 := 0; f1 <= 200/100; f1++ {
		for p50 := 0; p50 <= 200/50; p50++ {
			for p20 := 0; p20 <= 200/20; p20++ {
				for p10 := 0; p10 <= 200/10; p10++ {
					for p5 := 0; p5 <= 200/5; p5++ {
						for p2 := 0; p2 <= 200/2; p2++ {
							for p1 := 0; p1 <= 200/1; p1++ {
								sum = f1*100 + p50*50 + p20*20 + p10*10 + p5*5 + p2*2 + p1*1
								if sum == 200 {
									vaildCombiCount++
								}
							}
						}
					}
				}
			}
		}
	}
	return vaildCombiCount
}

func numberOfWaysToMakeX1(amt int) int {
	// defer timeTrack(time.Now(), "numberOfWaysToMakeX1()")

	// for 2 pounds + nothing else.
	vaildCombiCount := 1

	// follow the format of: for maxPosValue; minPosValue; --
	for f2 := 0; f2 <= amt/200; f2++ {
		rem1 := amt - f2*200
		for f1 := 0; f1 <= rem1/200; f1++ {
			rem2 := rem1 - f1*100
			for p50 := 0; p50 <= rem2/50; p50++ {
				rem3 := rem2 - p50*50
				for p20 := 0; p20 <= rem3/20; p20++ {
					rem4 := rem3 - p20*20
					for p10 := 0; p10 <= rem4/10; p10++ {
						rem5 := rem4 - p10*10
						for p5 := 0; p5 <= rem5/5; p5++ {
							rem6 := rem5 - p5*5
							for p2 := 0; p2 <= rem6/2; p2++ {
								// p1 := rem6 - p2*2
								// sum := f2*200 + f1*100 + p50*50 + p20*20 + p10*10 + p5*5 + p2*2 + p1*1
								// pf("f2:%d, f1:%d, p50:%d, p20:%d, p10:%d, p5:%d, p2:%d, p1:%d   sum = %d\n", f2, f1, p50, p20, p10, p5, p2, p1, sum)
								// if sum == amt {
								vaildCombiCount++
								// }
							}
						}
					}
				}
			}
		}
	}
	return vaildCombiCount
}

func numberOfWaysToMakeX2(amt int) int {
	// defer timeTrack(time.Now(), "numberOfWaysToMakeX1()")

	// for 2 pounds + nothing else.
	vaildCombiCount := 0

	// follow the format of: for maxPosValue; minPosValue; --
	for f2 := amt; f2 >= 0; f2 -= 200 {
		for f1 := f2; f1 >= 0; f1 -= 100 {
			for p50 := f1; p50 >= 0; p50 -= 50 {
				for p20 := p50; p20 >= 0; p20 -= 20 {
					for p10 := p20; p10 >= 0; p10 -= 10 {
						for p5 := p10; p5 >= 0; p5 -= 5 {
							for p2 := p5; p2 >= 0; p2 -= 2 {
								vaildCombiCount++
							}
						}
					}
				}
			}
		}
	}
	return vaildCombiCount
}

func main() {
	// p(numberOfWaysToMake2Pounds())
	p(numberOfWaysToMakeX1(200))
	p(numberOfWaysToMakeX2(200))

}
