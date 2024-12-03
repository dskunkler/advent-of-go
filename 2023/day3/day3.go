package day3

import (
	"advent/helper"
	"fmt"
	"regexp"
	"strconv"
)

func Part1() {
	textLines := helper.ArrayFromFileLines()
	getDigit := regexp.MustCompile(`(\d+)`)
  sol :=0
	for index, line := range textLines {
		nums := getDigit.FindAllStringIndex(line,-1)
		for _, digPos := range nums {
			fmt.Println(digPos)
			left, right := digPos[0], digPos[1]
			strVal, _ := strconv.Atoi(line[left:right]) 
			// check same line
			if helper.IsLeftRightOfLine(line, left, right){
				sol+= strVal
				fmt.Println("adding from left or right, ")
				continue
			}
			// check above
			if index > 0 {
				aboveLine := textLines[index-1]
				if helper.IsLeftRightOfLine(aboveLine, left, right){
					sol+= strVal
					fmt.Println("adding from lr above")
					continue
				}
				shouldContinue :=false
				for i:= left; i<right; i++ {
					if helper.IsNotNumberOrDot(aboveLine[i]) {
						fmt.Println("adding for above")
						sol+= strVal
						shouldContinue = true
						break
					}
				}
				if shouldContinue {
					continue
				}
			}
			// check below
			if index < len(textLines)-1 {
				belowLine := textLines[index+1]
				if helper.IsLeftRightOfLine(belowLine, left, right) {
					sol+= strVal
					fmt.Println("adding lr below")
					continue
				}
				shouldContinue:=false
				for i := left; i<right; i++ {
					if helper.IsNotNumberOrDot(belowLine[i]){
						fmt.Println("found below")
						shouldContinue = true
						sol+= strVal
						break
					}
				}
				if shouldContinue {
					continue
				}
			}
		}
	}

	fmt.Println(sol)
}

func Part2() {
	textLines := helper.ArrayFromFileLines()
	sol := 0
	for ind, line := range textLines {
		fmt.Println(ind, line)
		// Find star positions
		starRe := regexp.MustCompile(`\*`)
		starPos:= starRe.FindAllStringIndex(line,-1)
		if len(starPos) == 0 {
			// fmt.Println("no stars")
			continue
		}
		// fmt.Println(starPos)

		digRe := regexp.MustCompile(`(\d+)`)
		// Get surrounding digits
		currLineDigits := digRe.FindAllStringIndex(line,-1)
		aboveDigits := make([][]int,0)
		belowDigits := make([][]int,0)
		if ind > 0 {
			aboveDigits = digRe.FindAllStringIndex(textLines[ind-1],-1)
		}
		if ind < len(textLines)-1 {
			belowDigits = digRe.FindAllStringIndex(textLines[ind+1],-1)
		}

		// For each star, count neighbors and calculate ratio
		fmt.Println(starPos)
		for _, stars := range starPos {
			gearCt := 0
			gearRatio :=1
			star := stars[0]
			// Count next to the star
			fmt.Println(stars)
			for _, rng := range currLineDigits {
				left,right := rng[0], rng[1]
				if left == star+1 || right == star {
					fmt.Println("foundnext to")
					gearCt+=1
					tempNum,err := strconv.Atoi(line[left:right])
					if err == nil {
						gearRatio*=tempNum
					}
				}
			}
			if gearCt > 2 {
				continue
			}
			for _, rng:= range aboveDigits {
				left,right := rng[0], rng[1]
				if (left <= star+1 && left >= star-1 ) || (right >= star && right<=star+2) {
					fmt.Println("found above")
					gearCt+=1
					tempNum, err := strconv.Atoi(textLines[ind-1][left:right])
					if err == nil {
						gearRatio*=tempNum
					}
				}
			}
			if gearCt > 2 {
				continue
			}
			for _, rng:= range belowDigits {
				left,right := rng[0], rng[1]
				if (left <= star+1 && left >= star-1 ) || (right >= star && right<=star+2) {
					fmt.Println("found below")
					gearCt+=1
					tempNum, err := strconv.Atoi(textLines[ind+1][left:right])
					fmt.Println(textLines[ind+1])
					if err == nil {
						gearRatio*=tempNum
					}	
				}
			}
			if gearCt ==2 {
				sol+=gearRatio
				
				fmt.Println("2 gears, new sol", sol)
			}
		}
		

	}

	fmt.Println(sol)
}
