package day8

import (
	"advent/helper"
	"fmt"
)

func Part1() {
	textLines := helper.ArrayFromFileLines()
	directions := textLines[0]
	ourMap := make(map[string][]string)
	sol :=0
	// Make our map
	for i:= 2; i < len(textLines); i++ {
		line := textLines[i]
		key := line[:3]
		left := line[7:10]
		right := line[12:15]
		newAr := []string{left,right}
		ourMap[key] = newAr
	}
	currKey := "AAA"
	i:=0
	for currKey !="ZZZ" {
		sol+=1
		direction := directions[i]
		i = (i+1)%len(directions)
		if direction == 'L' {
			currKey = ourMap[currKey][0]
		} else {
			currKey = ourMap[currKey][1]
		}
	}

	fmt.Println(sol)
}


func Part2() {
	textLines := helper.ArrayFromFileLines()
	directions := textLines[0]
	ourMap := make(map[string][]string)
	// Make our map
	starts := make(map[string]int)
	for i:= 2; i < len(textLines); i++ {
		line := textLines[i]
		key := line[:3]
		if line[2] == 'A' {
			starts[key] = 1
		}
		left := line[7:10]
		right := line[12:15]

		newAr := []string{left,right}
		ourMap[key] = newAr
	}
	fmt.Println(starts)
	// for each start
	for start, _ := range starts {
		sol := 0
		i:=0
		currKey:=start	
		// search for first z
		for currKey[2] !='Z' {
			sol+=1
			direction := directions[i]
			i = (i+1)%len(directions)
			if direction == 'L' {
				currKey = ourMap[currKey][0]
				} else {
					currKey = ourMap[currKey][1]
				}
			}
			// take that number and assign it to start
			starts[start] = sol
	}

	// iterate through starts and find lcd
	fmt.Println(starts)
	results := make([]int,0)
	for _, val := range starts {
		results = append(results, val)
	}

	answer := results[0]
	for i :=1; i < len (results) ; i++ {
		fmt.Println("a", answer)
		fmt.Println("b",results[i])
		gcd := helper.GCD(results[i], answer)
		fmt.Println("gcd", gcd)
		answer *= results[i]/gcd
		fmt.Println("new answer", answer)
	}

}