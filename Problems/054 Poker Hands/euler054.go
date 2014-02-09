/*
In the card game poker, a hand consists of five cards and are ranked, from lowest to highest, in the following way:

High Card: Highest value card.
One Pair: Two cards of the same value.
Two Pairs: Two different pairs.
Three of a Kind: Three cards of the same value.
Straight: All cards are consecutive values.
Flush: All cards of the same suit.
Full House: Three of a kind and a pair.
Four of a Kind: Four cards of the same value.
Straight Flush: All cards are consecutive values of same suit.
Royal Flush: Ten, Jack, Queen, King, Ace, in same suit.
The cards are valued in the order:
2, 3, 4, 5, 6, 7, 8, 9, 10, Jack, Queen, King, Ace.

If two players have the same ranked hands then the rank made up of the highest
value wins; for example, a pair of eights beats a pair of fives (see example 1
below). But if two ranks tie, for example, both players have a pair of queens,
then highest cards in each hand are compared (see example 4 below); if the
highest cards tie then the next highest cards are compared, and so on.

The file, poker.txt, contains one-thousand random hands dealt to two players.
Each line of the file contains ten cards (separated by a single space): the
first five are Player 1's cards and the last five are Player 2's cards. You
can assume that all hands are valid (no invalid characters or repeated cards),
each player's hand is in no specific order, and in each hand there is a clear
winner.

How many hands does Player 1 win?
*/

package main

import (
	"fmt"
	"io/ioutil"
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

func functionName() {
	defer timeTrack(time.Now(), "functionName()")

}

func readFileAndProcess(fileName string) [][]string {

	fileBuf, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	// remove leading and trailing Unicode code points
	fileStr := strings.Trim(string(fileBuf), "")
	fileStr = strings.Trim(string(fileBuf), "\n")
	// split fileStr at each new line into a new string
	handSlice := strings.Split(fileStr, "\n")
	// initialise the array to the number of lines read in
	hands := make([][]string, len(handSlice))

	var line []string
	// for each base
	for i := range hands {
		line = strings.Split(handSlice[i], " ")
		hands[i] = make([]string, len(line))

		for j := range line {
			hands[i][j] = line[j]
		}
	}
	return hands
}

func cardValue(card string) int {
	valueMap := map[string]int{
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14}

	value := valueMap[card[:1]]

	if value != 0 {
		return value
	}
	return 0
}

func sortHandLowHigh(hand []string) []string {
	tempHand := append([]string(nil), hand...)
	sortedHand := []string{}
	for i := 0; i < len(tempHand); {
		//reset the lowestIdx and lowestCardValue
		lowestIdx := 0
		lowestCardValue := cardValue(tempHand[0])
		for j := 1; j < len(tempHand); j++ {
			if cardValue(tempHand[j]) < lowestCardValue {
				lowestIdx = j
				lowestCardValue = cardValue(tempHand[j])
			}
		}
		// add the lowest card value to the new storted list
		sortedHand = append(sortedHand, tempHand[lowestIdx])
		// remove the lowest card value from the list
		tempHand = append(tempHand[:lowestIdx], tempHand[lowestIdx+1:]...)
	}
	return sortedHand
}

func sortHandHighLow(hand []string) []string {
	tempHand := append([]string(nil), hand...)
	sortedHand := []string{}
	for i := 0; i < len(tempHand); {
		//reset the lowestIdx and lowestCardValue
		highestIdx := 0
		highestCardValue := cardValue(tempHand[0])
		for j := 1; j < len(tempHand); j++ {
			if cardValue(tempHand[j]) > highestCardValue {
				highestIdx = j
				highestCardValue = cardValue(tempHand[j])
			}
		}
		// add the highest card value to the new storted list
		sortedHand = append(sortedHand, tempHand[highestIdx])
		// remove the highest card value from the list
		tempHand = append(tempHand[:highestIdx], tempHand[highestIdx+1:]...)
	}
	return sortedHand
}

func sortValuesHighLow(values []int) []int {

	tempValues := append([]int(nil), values...)
	sortedValues := []int{}
	for i := 0; i < len(tempValues); {
		//reset the lowestIdx and lowestCardValue
		highestIdx := 0
		highestValue := tempValues[0]
		for j := 1; j < len(tempValues); j++ {
			if tempValues[j] > highestValue {
				highestIdx = j
				highestValue = tempValues[j]
			}
		}
		// add the highest value to the new storted list
		sortedValues = append(sortedValues, tempValues[highestIdx])
		// remove the highest card value from the list
		tempValues = append(tempValues[:highestIdx], tempValues[highestIdx+1:]...)
	}
	return sortedValues
}

func isFlush(hand []string) bool {
	// the suit from the first hand
	firstSuit := hand[0][1:2]
	for i := 1; i < len(hand); i++ {
		if hand[i][1:2] != firstSuit {
			return false
		}
	}
	return true
}

// isStraight returns true if the hand is a straight, and the value of the
// lowest card in the straight
func isStraight(hand []string) (bool, int) {
	sortedHand := sortHandLowHigh(hand)

	for i := 1; i < len(sortedHand); i++ {
		if cardValue(sortedHand[i]) != cardValue(sortedHand[i-1])+1 {
			return false, 0
		}
	}
	return true, cardValue(sortedHand[0])
}

// isFourOfAKind returns true if the hand has four of a kind, with a slice of
// the form [type, kicker]
func isFourOfAKind(hand []string) (bool, []int) {
	valueMap := map[int]int{}

	for i := 0; i < len(hand); i++ {
		valueMap[cardValue(hand[i])] += 1
	}
	result := []int{}
	for k, v := range valueMap {
		if v == 4 {
			// the type of the fours
			result = append(result, k)

			// find the kicker
			for k1, v1 := range valueMap {
				if v1 == 1 {
					result = append(result, k1)
				}
			}
			return true, result
		}
	}
	return false, result
}

func isFullHouse(hand []string) (bool, []int) {
	valueMap := map[int]int{}

	for i := 0; i < len(hand); i++ {
		valueMap[cardValue(hand[i])] += 1
	}
	twoPair := false
	twoPairVal := 0
	threePair := false
	threePairVal := 0

	result := []int{}
	for k, v := range valueMap {
		if v == 3 {
			// the type of the threes
			threePair = true
			threePairVal = k
		} else if v == 2 {
			// the type of the twos
			twoPair = true
			twoPairVal = k
		}
	}
	// if it's a full house
	if threePair && twoPair {
		result = append(result, threePairVal, twoPairVal)
		return true, result
	}
	return false, result
}

func isThreeOfAKind(hand []string) (bool, []int) {
	valueMap := map[int]int{}

	for i := 0; i < len(hand); i++ {
		valueMap[cardValue(hand[i])] += 1
	}
	threePair := false
	threePairVal := 0

	otherResults := []int{}

	result := []int{}
	for k, v := range valueMap {
		if v == 3 {
			// the type of the threes
			threePair = true
			threePairVal = k
		} else {
			otherResults = append(otherResults, k)
		}
	}
	// if it's a full house
	if threePair {
		result = append(result, threePairVal)

		if otherResults[0] > otherResults[1] {
			result = append(result, otherResults[0], otherResults[1])
		} else {
			result = append(result, otherResults[1], otherResults[0])
		}

		return true, result
	}
	return false, result
}

func isTwoPair(hand []string) (bool, []int) {
	valueMap := map[int]int{}

	for i := 0; i < len(hand); i++ {
		valueMap[cardValue(hand[i])] += 1
	}
	pairs := []int{}
	kicker := 0

	result := []int{}
	for k, v := range valueMap {
		if v == 2 {
			pairs = append(pairs, k)
		} else if v == 1 {
			kicker = k
		}
	}
	// if there are two pairs
	if len(pairs) == 2 {
		if pairs[0] > pairs[1] {
			result = append(result, pairs[0], pairs[1], kicker)
		} else {
			result = append(result, pairs[1], pairs[0], kicker)
		}
		return true, result
	}
	return false, result
}

func isOnePair(hand []string) (bool, []int) {
	valueMap := map[int]int{}

	for i := 0; i < len(hand); i++ {
		valueMap[cardValue(hand[i])] += 1
	}
	pair := 0
	kickers := []int{}

	result := []int{}
	for k, v := range valueMap {
		if v == 2 {
			pair = k
		} else if v == 1 {
			kickers = append(kickers, k)
		}
	}
	// if there is a pair
	if pair > 0 {
		result = append(result, pair)
		// sort the kickers
		kickers = sortValuesHighLow(kickers)
		result = append(result, kickers...)
		return true, result
	}
	return false, result
}

func makeSortedValues(hand []string) []int {
	sortedHand := sortHandHighLow(hand)
	values := []int{}

	for i := 0; i < len(sortedHand); i++ {
		values = append(values, cardValue(sortedHand[i]))
	}
	return values
}

func rankHand(hand []string) []int {

	// royal flush = [9]
	// straight flush = [8, lowcard value]
	// four of a kind = [7, value of 4 of a kind, kicker]
	// full house = [6, value of 3 of a kind, value of 2 of a kind]
	// flush = [5, kicker, kicker, kicker, kicker, kicker]
	// straight = [4, lowcard value]
	// three of a kind = [3, three of a kind value, kicker, kicker]
	// two pair = [2, highpair, lowpair, kicker]
	// one pair = [1, pair, kicker, kicker, kicker]
	// high card = [0, kicker, kicker, kicker, kicker, kicker]

	var check bool

	check, _ = isStraight(hand)
	if check {
		_, lowCardValue := isStraight(hand)
		// check for straight flush
		if isFlush(hand) {
			// check for royal flush
			if lowCardValue == 10 {
				return []int{9}
				// is a straight flush
			} else {
				return []int{8, lowCardValue}
			}
			// just a normal straight
		} else {
			return []int{4, lowCardValue}
		}
	}
	check, _ = isFourOfAKind(hand)
	if check {
		_, typeAndKicker := isFourOfAKind(hand)
		result := []int{7}
		result = append(result, typeAndKicker...)
		return result
	}
	check, _ = isFullHouse(hand)
	if check {
		_, score := isFullHouse(hand)
		result := []int{6}
		result = append(result, score...)
		return result
	}
	if isFlush(hand) {
		tempHand := sortHandHighLow(hand)
		result := []int{5}
		for i := 0; i < len(tempHand); i++ {
			result = append(result, cardValue(tempHand[i]))
		}
		return result
	}
	check, _ = isThreeOfAKind(hand)
	if check {
		result := []int{3}
		_, score := isThreeOfAKind(hand)
		result = append(result, score...)
		return result
	}
	check, _ = isTwoPair(hand)
	if check {
		result := []int{2}
		_, score := isTwoPair(hand)
		result = append(result, score...)
		return result
	}
	check, _ = isOnePair(hand)
	if check {
		result := []int{1}
		_, score := isOnePair(hand)
		result = append(result, score...)
		return result
	}
	result := []int{0}
	score := makeSortedValues(hand)
	result = append(result, score...)
	return result
}

func whoWinsHand(hands []string) int {
	rank1 := rankHand(hands[0:5])
	rank2 := rankHand(hands[5:10])

	for i := 0; ; i++ {
		// if both are the same check the next index
		if rank1[i] == rank2[i] {
			continue
			// player 1 wins
		} else if rank1[i] > rank2[i] {
			return 1
		}
		// player 2 wins
		return 2
	}
}

func gamesWonByPlayer1(hands [][]string) int {
	player1WinCount := 0

	for i := 0; i < len(hands); i++ {
		if whoWinsHand(hands[i]) == 1 {
			player1WinCount++
		}
	}
	return player1WinCount
}

func main() {
	defer timeTrack(time.Now(), "main()") // Timer function

	// royalflush := []string{"TC", "JC", "QC", "KC", "AC"}
	// straightflush := []string{"TC", "JC", "QC", "KC", "9C"}
	// flush := []string{"3C", "4C", "5C", "6C", "7C"}
	// straight := []string{"3C", "4S", "5C", "6C", "7C"}
	// fourkind := []string{"3C", "3S", "3D", "3H", "7C"}
	// fullhouse := []string{"3C", "3S", "3D", "8H", "8C"}
	// threekind := []string{"KC", "KS", "KD", "QH", "8C"}
	// twopair := []string{"8C", "KS", "QD", "QH", "8D"}
	// onepair := []string{"8C", "KS", "QD", "QH", "7D"}
	// highcard := []string{"8C", "KS", "QD", "JH", "7D"}

	// p(isFlush(flush))
	// p(isStraight(straight))
	// p(isFourOfAKind(fourkind))
	// p(isFullHouse(fullhouse))
	// p(isThreeOfAKind(threekind))
	// p(isTwoPair(twopair))
	// p(isOnePair(onepair))
	// p(makeSortedValues(highcard))
	// p("-------------------------")

	// p(rankHand(flush))
	// p(rankHand(royalflush))
	// p(rankHand(straightflush))
	// p(rankHand(onepair))
	// p(rankHand(highcard))
	// p("-------------------------")

	// p(sortValuesHighLow([]int{23, 44, 2, 7}))
	hands := readFileAndProcess("poker.txt")
	// p(sortHandLowHigh(hands[0][0:5]))
	// p(sortHandHighLow(hands[0][0:5]))

	// p("-------------------------")
	// p(whoWinsHand(hands[3]))

	p(gamesWonByPlayer1(hands))
	// p(hands)
}
