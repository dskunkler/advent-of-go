package day7

import (
	"advent/helper"
	"fmt"
	"sort"
	"strconv"
)
var order = "AKQT98765432J"

func indexInOrder(char byte) int {
	for i, v := range order {
		if byte(v) == char {
			return i
		}
	}
	return -1
}

func customSort(strings []string) {
	sort.Slice(strings, func(i, j int) bool {
		str1, str2 := strings[i], strings[j]

		for k := 0; k < 5; k++ {
			index1 := indexInOrder(str1[k])
			index2 := indexInOrder(str2[k])

			if index1 > index2 {
				return true
			} else if index1 < index2 {
				return false
			}
		}

		return false
	})
}

func Part1() {
	sol := 0
	textLines := helper.ArrayFromFileLines()
	fiveOfAKind:= make([]string, 0)
	fourOfAKind:= make([]string, 0)
	fullHouse:= make([]string, 0)
	threeOfAKind:= make([]string, 0)
	twoPair:= make([]string, 0)
	onePair:= make([]string, 0)
	highCard:= make([]string, 0)

	// Organize into groups
	for _, line := range textLines { 
		hand := line[:5]
		tempMap := make(map[string]int)
		for i := 0; i < 5; i++ {
			tempMap[string(hand[i])]+=1
		}
		if len(tempMap) == 1  {
			fiveOfAKind = append(fiveOfAKind, line)
		} else if len(tempMap) == 5 {
			highCard = append(highCard, line)
		} else if len(tempMap) == 4 {
			onePair = append(onePair, line)
		} else if len(tempMap) == 2 {
			for _, val :=range tempMap {
				if val == 4 {
					fourOfAKind = append(fourOfAKind, line)
					break
				} else if val == 3 {
					fullHouse = append(fullHouse, line)
					break
				}
			}
		} else {
			for _, val := range tempMap {
				if val == 3 {
					threeOfAKind = append(threeOfAKind, line)
					break
				} else if val == 2 {
					twoPair = append(twoPair, line)
					break
				}
			}
		}
	}
	// organize groups lowest to smallest
	// do the math
	counter :=0
	customSort(highCard)
	fmt.Println(highCard)
	for _,line := range highCard{
		counter+=1
		num, _ := strconv.Atoi(line[6:])
		sol += counter * num

	}
	customSort(onePair)
	fmt.Println(onePair)
	for _,line := range onePair{
		counter+=1
		num, _ := strconv.Atoi(line[6:])
		sol += counter * num
	}	
	customSort(twoPair)
	fmt.Println(twoPair)
	for _,line := range twoPair{
		counter+=1
		num, _ := strconv.Atoi(line[6:])
		sol += counter * num
	}
	customSort(threeOfAKind)
	fmt.Println(threeOfAKind)
	for _,line := range threeOfAKind{
		counter+=1
		num, _ := strconv.Atoi(line[6:])
		sol += counter * num
	}
	
customSort(fullHouse)
fmt.Println(fullHouse)
	for _,line := range fullHouse{
		counter+=1
		num, _ := strconv.Atoi(line[6:])
		sol += counter * num
	}
	customSort(fourOfAKind)
	fmt.Println(fourOfAKind)
	for _,line := range fourOfAKind{
		counter+=1
		num, _ := strconv.Atoi(line[6:])
		sol += counter * num
	}
	customSort(fiveOfAKind)
	fmt.Println(fiveOfAKind)
	for _,line := range fiveOfAKind{
		counter+=1
		num, _ := strconv.Atoi(line[6:])
		sol += counter * num
	}
	
	fmt.Println(sol)


}

func Part2() {
	sol := 0
	textLines := helper.ArrayFromFileLines()
	fiveOfAKind:= make([]string, 0)
	fourOfAKind:= make([]string, 0)
	fullHouse:= make([]string, 0)
	threeOfAKind:= make([]string, 0)
	twoPair:= make([]string, 0)
	onePair:= make([]string, 0)
	highCard:= make([]string, 0)
	fmt.Println("~~~len", len(textLines))
	// Organize into groups
	for _, line := range textLines { 
		hand := line[:5]
		tempMap := make(map[string]int)
		wilds:=0
		for i := 0; i < 5; i++ {
			card:= string(hand[i])
			if card == "J" {
				wilds+=1
			} else {
				tempMap[card]+=1
			}
		}
		if len(tempMap) == 1  || len(tempMap) == 0 {
			// all same kind
			fiveOfAKind = append(fiveOfAKind, line)
		} else if len(tempMap) == 5 {
			// All unique
			highCard = append(highCard, line)
		} else if len(tempMap) == 4 {
			// If theres a wild card, then we have a pair +1 so three
			// If theres 4, it means we'll ALWAYS have one pair
			onePair = append(onePair, line)
		} else if len(tempMap) == 2 {
			// could be lots of variations here....
			// 0 wilds regular
			// 1 wild, (3,1) , (2,2)
			// 2 wilds (1, 2) really only one option here
			// 3 wilds ( 1,1) only one option
			for _, val :=range tempMap {
				if wilds == 0 {
					if val == 4 {
						fourOfAKind = append(fourOfAKind, line)
						break
						// could be three + 2 now
					} else if val == 3 {
						fullHouse = append(fullHouse, line)
						break
					}
				} else if wilds == 1 {
					if val == 3 {
						fourOfAKind = append(fourOfAKind, line)
						break
					} else if val == 2 {
						fullHouse = append(fullHouse, line)
						break
					}
				} else { // 2 and 3
					fourOfAKind = append(fourOfAKind, line)
					break
				}
			}
		} else { // len 3
			// 0 wilds (3,1,1) (2,2,1)
			// 1 wild (2, 1,1)
			// 2 wlids ( 1,1,1)
			for _, val := range tempMap {
				if wilds > 0 {
					threeOfAKind = append(threeOfAKind, line)
					break
				} else {
					if val == 3 {
						threeOfAKind = append(threeOfAKind, line)
						break
					} else if val == 2 {
						twoPair = append(twoPair, line)
						break
					}
				}
			}
		}
	}
	// organize groups lowest to smallest
	// do the math

	counter :=0
	customSort(highCard)
	// fmt.Println("high",highCard)
	for _,line := range highCard{
		counter+=1
		num, _ := strconv.Atoi(line[6:])
		sol += counter * num

	}
	customSort(onePair)
	// fmt.Println("pair",onePair)
	for _,line := range onePair{
		fmt.Println(line[:5])
		counter+=1
		num, _ := strconv.Atoi(line[6:])
		sol += counter * num

	}	
	customSort(twoPair)
	// fmt.Println("twopair",twoPair)
	for _,line := range twoPair{
		counter+=1
		num, _ := strconv.Atoi(line[6:])
		// sol += counter * num
		sol += counter * num

	}
	customSort(threeOfAKind)
	// fmt.Println("three",threeOfAKind)
	for _,line := range threeOfAKind{
		counter+=1
		num, _ := strconv.Atoi(line[6:])
		// sol += counter * num
		sol += counter * num

	}
	
customSort(fullHouse)
// fmt.Println("full",fullHouse)
	for _,line := range fullHouse{
		counter+=1
		num, _ := strconv.Atoi(line[6:])
		// sol += counter * num
		sol += counter * num

	}
	customSort(fourOfAKind)
	// fmt.Println("four",fourOfAKind)
	for _,line := range fourOfAKind{
		counter+=1
		num, _ := strconv.Atoi(line[6:])
		// sol += counter * num
		sol += counter * num

	}
	customSort(fiveOfAKind)
	// fmt.Println("five",fiveOfAKind)
	for _,line := range fiveOfAKind{
		counter+=1
		num, _ := strconv.Atoi(line[6:])
		// sol += counter * num
		sol += counter * num

	}
	fmt.Println(sol)

}