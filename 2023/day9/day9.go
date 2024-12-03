package day9

import (
	"advent/helper"
	"fmt"
	"os"
)

func constructLine(line []int ) []int {
	tempDiff := make([]int,0)
	isAllZeros := true
	fmt.Println("running on", line)
	// check for all zeros base case
	for _,num := range line {
		if num != 0 {
			isAllZeros = false
			break
		}
	}
	if isAllZeros {
		fmt.Println("is all zeros")
		return line
	}
	fmt.Println("not all zeros")
	// Make the diff array
	for i :=0; i < len(line)-1; i++ {
		// create the diff
		tempDiff = append(tempDiff,line[i+1] - line[i])
	}
	fmt.Println("diff array", tempDiff)
	// Generate next line
	tempDiff = constructLine(tempDiff)
	fmt.Println("constructed dif", tempDiff)

	line = append(line, line[len(line)-1] + tempDiff[len(tempDiff)-1])
	fmt.Println("generated fine line", line)
	return line
}

func constructReverseLine(line []int ) []int {
	tempDiff := make([]int,0)
	isAllZeros := true
	// check for all zeros base case
	for _,num := range line {
		if num != 0 {
			isAllZeros = false
			break
		}
	}
	if isAllZeros {
		return line
	}
	// Make the diff array
	for i :=0; i < len(line)-1; i++ {
		// create the diff
		tempDiff = append(tempDiff,line[i+1] - line[i])
	}
	// Generate next line
	tempDiff = constructReverseLine(tempDiff)
	line = append([]int{line[0]- tempDiff[0]},line...)
	return line
}
func Part1() {
	textLines := helper.ArrayFromFileLines()
	fmt.Println(textLines)
	sol:=0
	for _, line := range textLines {
		lineDigits, err := helper.NumberArrayFromString(line)
		if err != nil {
			fmt.Println("non digit in string")
			os.Exit(1)
		}
		res := constructLine(lineDigits)
	if len(res) == 0 {
		sol+=0
	} else {
		sol+=res[len(res)-1]
	}	
		fmt.Println("res", res)
	fmt.Println("new sol",sol)	
	}

	fmt.Println(sol)
}


func Part2() {
	textLines := helper.ArrayFromFileLines()
	fmt.Println(textLines)
	sol:=0
	for _, line := range textLines {
		lineDigits, err := helper.NumberArrayFromString(line)
		if err != nil {
			fmt.Println("non digit in string")
			os.Exit(1)
		}
		res := constructReverseLine(lineDigits)
		if len(res) == 0 {
			sol+=0
		} else {
			sol+=res[0]
		}	
		fmt.Println("res", res)
		fmt.Println("new sol",sol)
	}

	fmt.Println(sol)
}