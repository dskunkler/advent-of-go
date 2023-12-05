package day2

import (
	"advent/helper"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day2() {
	textLines := helper.ArrayFromFileLines()
	sol:=0
	red:= regexp.MustCompile(`(\d+) red`)
	blue:= regexp.MustCompile(`(\d+) blue`)
	green:= regexp.MustCompile(`(\d+) green`)
	for _, line := range textLines {
		re := regexp.MustCompile(`Game (\d+):`)
		id := re.FindStringSubmatch(line)
		temp := strings.Split(line, ":")
		parts := strings.Split(temp[1], ";")
		fmt.Println(parts)
		idVal, _ := strconv.Atoi(id[1])
		fmt.Println("adding", idVal)
			sol+=idVal
		for _, picks :=range parts {
			reds:= red.FindStringSubmatch(picks)
			blues:= blue.FindStringSubmatch(picks)
			greens:= green.FindStringSubmatch(picks)
			fmt.Println(picks)	
			if len(reds) >= 2 {
				numReds,_ := strconv.Atoi(reds[1])
				fmt.Println("numsReds", numReds)
				if numReds > 12{
					fmt.Println(numReds, "oob")
					sol-=idVal
					break
				}
			}	
			if len(blues) >= 2 {
				numBlues,_ := strconv.Atoi(blues[1])
				if numBlues > 14{
					fmt.Println( "oob")
					sol-=idVal
					break
				}
			}	
			if len(greens) >= 2 {
				numGreens,_ := strconv.Atoi(greens[1])
				if numGreens > 13{
					fmt.Println( "oob")
					sol-=idVal
					break
				}
			}	
			
			fmt.Println("new sol", sol)
		}
	}
	fmt.Println(sol)
}

func Part2() {
	textLines := helper.ArrayFromFileLines()
	sol:=0
	red:= regexp.MustCompile(`(\d+) red`)
	blue:= regexp.MustCompile(`(\d+) blue`)
	green:= regexp.MustCompile(`(\d+) green`)
	for _, line := range textLines {
		temp := strings.Split(line, ":")
		parts := strings.Split(temp[1], ";")
		fmt.Println(parts)
		redMax := 0
		blueMax := 0
		greenMax := 0
		for _, picks :=range parts {
			// Find max for each in picks
			reds:= red.FindStringSubmatch(picks)
			blues:= blue.FindStringSubmatch(picks)
			greens:= green.FindStringSubmatch(picks)
			fmt.Println(picks)	
			if len(reds) >= 2 {
				numReds,_ := strconv.Atoi(reds[1])
				redMax = helper.Max(redMax, numReds)
			}	
			if len(blues) >=2 {
				numBlues,_ := strconv.Atoi(blues[1])
				blueMax = helper.Max(blueMax, numBlues)
			}
			if len(greens) >=2 {
				numGreens,_ :=strconv.Atoi(greens[1])
				greenMax = helper.Max(greenMax, numGreens)
			}
		}
		fmt.Println("red", redMax, " blue", blueMax, " green", greenMax)
		power := redMax*blueMax *greenMax
		fmt.Println("power: ",power)
		sol+=power
		// multiply maxes and add to sol
	}
	fmt.Println(sol)
}