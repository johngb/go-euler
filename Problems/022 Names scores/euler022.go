package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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

// convert a char to an int.  No checking int place
func alphabeticValue(c byte) int {
	return int(c - 'A' + 1)
}

func nameLetterSum(name string) int {
	sum := 0
	for i := 0; i < len(name); i++ {
		sum += alphabeticValue(name[i])
	}
	return sum
}

func letterHash(fileName string) int {
	defer timeTrack(time.Now(), "letterHash()") // Timer function

	fileBuf, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	// remove leading and trailing Unicode code points
	fileStr := strings.Trim(string(fileBuf), "")
	// remove all quotation marks (")
	fileStr = strings.Replace(fileStr, "\"", "", -1)
	nameList := strings.Split(fileStr, ",")
	sort.Strings(nameList)

	sum := 0
	for i := 0; i < len(nameList); i++ {
		sum += nameLetterSum(nameList[i]) * (i + 1)
	}
	return sum
}

func main() {

	p(letterHash("names.txt"))

}
