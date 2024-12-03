package main

import (
	"advent/helper"
	"fmt"
	"strings"
)

func getPossibilitiesInLine(line string) int {
	spacePos := strings.Index(line, " " )
	record := line[:spacePos]
	errors := line[spacePos+1:]

	fmt.Println(record)
	fmt.Println(errors)

	return -1

}

func part1() {
	// Observations.
	// we always need a seperator so for n errors, we need n + n-1 spaces
	// if we have an exact amount of broken, we can remove it from our list as it must be that
	stringLines := helper.ArrayFromFileLines()
	for _,line :=range stringLines {
		getPossibilitiesInLine(line)
	} 



}

func main() {
	part1()
}