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

func primesNoSieve(max int) []int {
	defer timeTrack(time.Now(), "listOfPrimes()") // Timer function

	primes := []int{2}
	// next to check = 1 + the last number in the primes list
	pPrime := 3
	// iterate through each prime value in the known primes list
	for ; pPrime < max; pPrime += 2 {
		maxFactor := int(math.Sqrt(float64(pPrime)))

		// idx = 1, as we never need to check the first prime (i.e. 2) because
		// we start with an odd number and increment by 2
		for idx := 1; idx < len(primes); idx++ {
			// if remaining primes are bigger than maximum possible factor size
			if primes[idx] > maxFactor {
				// pPrime must be a prime, so can end loop
				break
			}
			if pPrime%primes[idx] == 0 {
				goto newloop
			}
		}
		// pPrime is a prime, so add to list of primes
		primes = append(primes, pPrime)
	newloop:
	}
	return primes
}

func primeSieve(max int) []int {
	// defer timeTrack(time.Now(), "primeSieve") // Timer function

	sieve := make([]bool, max/2+1)
	// initialise the prime list with the only even prime
	primeList := []int{2}

	// 3 is the first odd prime
	for i := 3; i <= max; i += 2 {
		if sieve[i/2] == false {
			// i should be a prime number, so add to list of primes
			primeList = append(primeList, i)
			// any odd * even = even, so avoid all even multiples
			for j := 3; i*j <= max; j += 2 {
				sieve[i*j/2] = true
			}
		}
	}
	return primeList
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

func isPermutation(a, b int) bool {

	strA := strconv.Itoa(a)
	strB := strconv.Itoa(b)

	for i := 0; i < len(strA); {
		char := string(strA[0])
		if strings.Contains(strB, char) {
			//remove that character once from both strings
			strA = strings.Replace(strA, char, "", 1)
			strB = strings.Replace(strB, char, "", 1)
		} else {
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
