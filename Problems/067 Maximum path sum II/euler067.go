package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
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

	fileBuf, err := ioutil.ReadFile("euler067_data.txt")
	if err != nil {
		panic(err)
	}
	// remove leading and trailing Unicode code points
	fileStr := strings.Trim(string(fileBuf), "")
	// split fileStr at each new line into a new string
	oneDArrStr := strings.Split(fileStr, "\n")
	var line []string
	// initialise the array to the number of lines read in
	arr := make([][]int, len(oneDArrStr))

	for i := range arr {
		line = strings.Split(oneDArrStr[i], " ")
		arr[i] = make([]int, len(line))

		for j := range line {
			arr[i][j], _ = strconv.Atoi(line[j])
		}
	}

	p("Max path: ", maxPathSum(arr))
	// p(arr)
}
