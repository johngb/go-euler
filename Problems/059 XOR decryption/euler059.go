/*
Each character on a computer is assigned a unique code and the preferred
standard is ASCII (American Standard Code for Information Interchange). For
example, uppercase A = 65, asterisk (*) = 42, and lowercase k = 107.

A modern encryption method is to take a text file, convert the bytes to ASCII,
then XOR each byte with a given value, taken from a secret key. The advantage
with the XOR function is that using the same encryption key on the cipher
text, restores the plain text; for example, 65 XOR 42 = 107, then 107 XOR 42 =
65.

For unbreakable encryption, the key is the same length as the plain text
message, and the key is made up of random bytes. The user would keep the
encrypted message and the encryption key in different locations, and without
both "halves", it is impossible to decrypt the message.

Unfortunately, this method is impractical for most users, so the modified
method is to use a password as a key. If the password is shorter than the
message, which is likely, the key is repeated cyclically throughout the
message. The balance for this method is using a sufficiently long password key
for security, but short enough to be memorable.

Your task has been made easy, as the encryption key consists of three lower
case characters. Using cipher1.txt (right click and 'Save Link/Target As...'),
a file containing the encrypted ASCII codes, and the knowledge that the plain
text must contain common English words, decrypt the message and find the sum
of the ASCII values in the original text.
*/

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

// timeTrack is used for basic benchmarking in other functions
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}

func readTextFileAndProcess(fileName string) []int {
	// defer timeTrack(time.Now(), "readTextFileAndProcess()") // Timer function

	fileBuf, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	// remove leading and trailing Unicode code points
	fileStr := strings.Trim(string(fileBuf), "\n")
	// split fileStr at each new line into a new string
	strSlice := strings.Split(fileStr, ",")

	arr := make([]int, len(strSlice))

	for i := 0; i < len(strSlice); i++ {
		arr[i], _ = strconv.Atoi(strSlice[i])
	}

	return arr
}

func decodeUsingKey(codedFile []int, key []int) []int {
	xorFile := make([]int, len(codedFile))
	for idx := 0; idx < len(codedFile); idx++ {
		keyidx := idx % len(key)
		xorFile[idx] = codedFile[idx] ^ key[keyidx]
	}
	return xorFile
}

func decodeUsingKey1(codedFile []int, key []int, portionToCheck int) []int {
	length := len(codedFile) / portionToCheck
	xorFile := make([]int, length)
	for idx := 0; idx < length; idx++ {
		keyidx := idx % len(key)
		xorFile[idx] = codedFile[idx] ^ key[keyidx]
	}
	return xorFile
}

func convertToString(file []int) string {
	str := ""
	for i := 0; i < len(file); i++ {
		str = str + string(file[i])
	}
	return str
}

func charSum(data []int, charString string) int {
	sum := 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(charString); j++ {
			if data[i] == int(charString[j]) {
				sum++
				break
			}
		}
	}
	return sum
}

func decryptFile(fileName string, wordToFind string) int {
	defer timeTrack(time.Now(), "decryptFile()")

	intCodedFile := readTextFileAndProcess(fileName)
	tempFile := make([]int, len(intCodedFile))

	key := []int{0, 0, 0}
	a, b, c := 0, 0, 0
	tempString := ""

	//97 is the ascii for 'a', and 122 is the ascii for 'z'
	for a = 97; a <= 122; a++ {
		for b = 97; b <= 122; b++ {
			for c = 97; c <= 122; c++ {
				key[0] = a
				key[1] = b
				key[2] = c
				tempFile = decodeUsingKey(intCodedFile, key)
				// CHECK FOR VALID STRING

				tempString = convertToString(tempFile)
				if strings.Contains(tempString, wordToFind) {
					goto finishedloop
				}
			}
		}
	}
finishedloop:
	// p("key was: ", key)
	count := 0
	for i := 0; i < len(tempFile); i++ {
		count += tempFile[i]
	}
	return count
}

func decryptFile1(fileName string, charString string, portionToCheck int) int {
	defer timeTrack(time.Now(), "decryptFile1()")

	intCodedFile := readTextFileAndProcess(fileName)
	tempFile := make([]int, len(intCodedFile)/portionToCheck)

	key := []int{0, 0, 0}
	a, b, c := 0, 0, 0
	maxCharSum := 0
	maxCharSumKey := []int{0, 0, 0}

	//97 is the ascii for 'a', and 122 is the ascii for 'z'
	for a = 97; a <= 122; a++ {
		for b = 97; b <= 122; b++ {
			for c = 97; c <= 122; c++ {
				key[0] = a
				key[1] = b
				key[2] = c
				tempFile = decodeUsingKey1(intCodedFile, key, portionToCheck)
				// CHECK FOR VALID STRING
				sumOfChars := charSum(tempFile, charString)
				if sumOfChars > maxCharSum {
					maxCharSum = sumOfChars
					copy(maxCharSumKey, key)
				}
			}
		}
	}

	// use the key with the highest char count and perform for the entire file
	completeFile := decodeUsingKey1(intCodedFile, maxCharSumKey, 1)
	count := 0
	for i := 0; i < len(completeFile); i++ {
		count += completeFile[i]
	}
	return count
}

func main() {
	// str := "aztgtt"
	// p(str[0])
	// str = str + string(112)
	// p(str)

	// arr := readTextFileAndProcess("cipher1.txt")
	// p(arr[0] ^ int(str[2]))
	// p(str[0], str[1])

	word := "world"
	p(decryptFile("cipher1.txt", word))

	p(decryptFile1("cipher1.txt", "ent s", 40))

}
