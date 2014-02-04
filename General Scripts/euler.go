package euler

import (
	"fmt"
	"io/ioutil"
	"math/big"
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

func Factorial(n int64) *big.Int {
	if n < 0 {
		return big.NewInt(1)
	}
	if n == 0 {
		return big.NewInt(1)
	}
	bigN := big.NewInt(n)
	return bigN.Mul(bigN, factorial(n-1))
}

func BigFactorial(n int64) *big.Int {
	if n < 0 {
		return big.NewInt(1)
	}
	if n == 0 {
		return big.NewInt(1)
	}
	bigN := big.NewInt(n)
	return bigN.Mul(bigN, factorial(n-1))
}

func mapOfPrimes(min, max int) map[int]bool {
	// defer timeTrack(time.Now(), "mapOfPrimes") // Timer function

	sieve := make([]bool, max+2)
	primeMap := make(map[int]bool)

	// 2 is the first prime
	for i := 2; i <= max; i++ {
		if sieve[i] == false {
			// i should be a prime number, so add to list of primes
			primeMap[i] = true
			for j := 2; i*j <= max; j++ {
				sieve[i*j] = true
			}
		}
	}
	return primeMap
}

func isPandigital(a int) bool {
	// defer timeTrack(time.Now(), "isPandigital()")

	str := strconv.Itoa(a)
	if len(str) != 9 {
		return false
	}

	for i := 1; i <= 9; i++ {
		// if the string doesn't contain one of the numbers
		if !strings.Contains(str, strconv.Itoa(i)) {
			return false
		}
	}
	return true
}

// isPrime returns true if 'pPrime' is a prime
func isPrime(pPrime int) bool {
	if pPrime < 0 {
		return false
	}
	maxFactor := int(math.Sqrt(float64(pPrime)))
	for i := 2; i <= maxFactor; i++ {
		if pPrime%i == 0 {
			return false
		}
	}
	return true
}

// readTextFileAndProcess reads in a text file with comma separated quoted
// words and returns an unsorted slice of strings where each string is a
// single word
func readTextFileAndProcess(fileName string) []string {
	defer timeTrack(time.Now(), "readTextFileAndProcess()") // Timer function

	fileBuf, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	// remove leading and trailing Unicode code points
	fileStr := strings.Trim(string(fileBuf), "")
	// remove all quotation marks (")
	fileStr = strings.Replace(fileStr, "\"", "", -1)
	// remove all newlines
	fileStr = strings.Replace(fileStr, "\n", "", -1)
	nameList := strings.Split(fileStr, ",")

	return nameList
}
