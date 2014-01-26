package main

import (
	"fmt"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

// Used to time other functions to compare speed
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

func minPathSum(mx [][]int) int {
	defer timeTrack(time.Now(), "minPathSum()")

	a, b := 0, 0
	for row := len(mx) - 2; row >= 0; row-- {
		for col := len(mx[row]) - 1; col >= 0; col-- {
			a = mx[row+1][col]
			b = mx[row+1][col+1]
			if a < b {
				mx[row][col] += a
			} else {
				mx[row][col] += b
			}
		}
	}
	return mx[0][0]
}

func maxPathSum(mx [][]int) int {
	defer timeTrack(time.Now(), "maxPathSum()")

	a, b := 0, 0
	for row := len(mx) - 2; row >= 0; row-- {
		for col := len(mx[row]) - 1; col >= 0; col-- {
			a = mx[row+1][col]
			b = mx[row+1][col+1]
			if a > b {
				mx[row][col] += a
			} else {
				mx[row][col] += b
			}
		}
	}
	return mx[0][0]
}

func main() {

	data := [][]int{{75},
		{95, 64},
		{17, 47, 82},
		{18, 35, 87, 10},
		{20, 4, 82, 47, 65},
		{19, 1, 23, 75, 3, 34},
		{88, 2, 77, 73, 7, 63, 67},
		{99, 65, 4, 28, 6, 16, 70, 92},
		{41, 41, 26, 56, 83, 40, 80, 70, 33},
		{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
		{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
		{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
		{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
		{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
		{04, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23}}

	// testData := [][]int{{3},
	// 	{7, 4},
	// 	{2, 4, 6},
	// 	{8, 5, 9, 3}}

	// p("Min path: ", minPathSum(data))
	p("Max path: ", maxPathSum(data))
	// p("Max path test: ", maxPathSum(testData))

}
