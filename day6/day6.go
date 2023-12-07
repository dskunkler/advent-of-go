package day6

import (
	"advent/helper"
	"fmt"
	"regexp"
	"strconv"
)

func Part1() {
	textLines := helper.ArrayFromFileLines()
	digRe := regexp.MustCompile(`(\d+)`)
	fmt.Println(textLines)
	times := digRe.FindAllString(textLines[0],-1)
	distances := digRe.FindAllString(textLines[1],-1)
	sol :=1
	for i, time := range times {
		distance := distances[i]
		tempsols :=0
		timeNum, _:= strconv.Atoi(time)
		distNumber, _ := strconv.Atoi(distance)
		fmt.Println(timeNum, distNumber)
		// number of seconds waited * time - seconds
		for speed :=1; speed < timeNum; speed++ {
			if speed * (timeNum -speed) > distNumber {
				tempsols++
			} 
		}
		// fmt.Println("num found", tempsols)
		sol*=tempsols
	}
	fmt.Println(sol)
}
