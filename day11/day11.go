package day11

import (
	"advent/helper"
	"fmt"
	"math"
)
func getDistance(galaxies [][2]int) int {
	result:=0
	for i, galaxy := range galaxies {
		for _, otherGalaxy := range galaxies[i+1:] {
			distance :=int( math.Abs(float64(galaxy[0] - otherGalaxy[0])) +math.Abs(float64((galaxy[1] -otherGalaxy[1]) )))
			result+=distance
		}
	}
	return result	
}
func getGalaxyByOffset(byteLines [][]byte, offset int)[][2]int{
	galaxyPositions := make([][2]int,0)
	rowOffset:= make([]int,0)
	currOffset:= 0
	for i :=0 ; i < len(byteLines)  ; i+=1 {
		sawNumber:=false
		for j :=0 ; j < len(byteLines[0]); j+=1 {
			// if we see an octothorp, number it
			if byteLines[i][j] == '#' {
				sawNumber = true
				break
			}
		}
		rowOffset = append(rowOffset, currOffset)

		if !sawNumber {
			currOffset+=offset
		}
	}
	currOffset = 0
	for j :=0; j < len(byteLines[0]); j++ {
		sawNumber := false
		for i :=0; i < len(byteLines); i++ {
			if byteLines[i][j] != '.' {
				fmt.Println("found at ",i,j)
				galaxyPositions = append(galaxyPositions, [2]int{i + rowOffset[i], j + currOffset})
				fmt.Println(galaxyPositions)
				sawNumber = true
			}
		}
		if !sawNumber {
			currOffset+=offset
		}
	}
	return galaxyPositions
}

func Part1() {
	// go through file and find all rows and columns that need to be doubled

	stringLines := helper.ArrayFromFileLines()
	byteLines := helper.BytesFromStringArray(stringLines)
	galaxy := getGalaxyByOffset(byteLines, 1)
	fmt.Println(galaxy)
	// helper.PrintByteArray(byteLines)
	// myMap := make(map[byte]int)
	// go through every galaxy
	sol:=0
	sol = getDistance(galaxy)

	fmt.Println(sol)
}

func Part2() {
	stringLines := helper.ArrayFromFileLines()
	byteLines := helper.BytesFromStringArray(stringLines)
	galaxy := getGalaxyByOffset(byteLines, 1000000-1)
	fmt.Println(galaxy)
	fmt.Println(getDistance(galaxy))
}