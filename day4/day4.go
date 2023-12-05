package day4

import (
	"advent/helper"
	"fmt"
	"math"
	"strings"
)

func Part1() {
	textLines := helper.ArrayFromFileLines()
	// fmt.Println(textLines)
	total :=0.0
	for _, line:= range textLines {
		parts := strings.Split(line, ":") 
		numbers := parts[1]
		halves := strings.Split(numbers,"|" )
		// fmt.Println(halves[0])
		set := make(map[string]int)
		sols := strings.Split(halves[0], " ")
		fmt.Println(sols)
		for _, num := range sols[1:] {
			// fmt.Println(num)
			set[num] = 1
		}
		found := 0.0
		score := 0.0
		delete(set, "")
		fmt.Println(set)
		possible := strings.Split(halves[1], " ")
		for _, num := range possible {
			// fmt.Println(num)
			
			if _, ok := set[num]; ok {
				fmt.Println("ok", num)
				score = math.Pow(2, found)
				found+=1
			}
		}
		fmt.Println("score: ",score)
		total += score
	}
	fmt.Println(total)
}

func Part2() {
	textLines := helper.ArrayFromFileLines()
	// fmt.Println(textLines)
	sol :=0
	copies := make(map[int]int)
	// go over lines
	for ind, line := range textLines {
		fmt.Println(line)
		// Set copies as initial one if needed
		if _,ok:= copies[ind];!ok {
			// fmt.Println("new thing adding key", id[2])
			copies[ind] = 1
		}

		// Split the line
		parts := strings.Split(line, ":") 
		numbers := parts[1]
		halves := strings.Split(numbers,"|" )

		// Make a set for the lines solutions
		set := make(map[string]int)
		sols := strings.Split(halves[0], " ")
		for _, num := range sols[1:] {
			set[num] = 1
		}
		delete(set, "")
		possible := strings.Split(halves[1], " ")
		
		// Find the number of solutions from the card
		found := 0
		for _, num := range possible {
			if _, ok := set[num]; ok {
				fmt.Println("ok", num)
				found+=1
			}
		}
		fmt.Println("found",found)
		// Add copies for next cards
		for i:=1; i <= found; i++ {
			// Add the current number of copies to that card
			// since every card from this index will go to the found ranged one
			fmt.Println("i",i)
			 x,ok :=copies[ind+i]
			fmt.Println("x=", x,ok) 
			if ok {
				copies[ind+i] += copies[ind]
			} else {
				copies[ind+i] = 1+copies[ind]
			}
			fmt.Println(copies)
		}
		fmt.Println("copies now", copies)
		// Add number of copies we have now to solution
		sol+=copies[ind]
	}

	
	fmt.Println(sol)
}