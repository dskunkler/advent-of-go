package main

import (
	"advent/helper"
	"fmt"
)

func ShiftZerosNorth(lines [][]byte) [][]byte {
	for i:=0; i < len(lines)-1; i++ {
		for j := 0; j < len(lines[0]); j++ {
			if lines[i][j] != byte('.') {
				continue
			}
			for x:=i+1; x < len(lines); x++ {
				if lines[x][j] == byte('#') {
					break
				}
				if lines[x][j] == byte('.') {
					continue
				}
				lines[i][j] = byte('O')
				lines[x][j] = byte('.')
				break
			}
		}
	}
	return lines
}

func ShiftZerosWest(lines [][]byte) [][]byte {
	for j:= 0; j < len(lines[0])-1; j++ {
		for i :=0; i < len(lines); i++ {
			if lines[i][j] != byte('.'){
				continue
			}
			// fmt.Println("starting at i=", i, "j=", j)
			for x:=j+1; x< len(lines[0]); x++ {
				if lines[i][x] == byte('#') {
					// fmt.Println("hit blocker")
					break
				}
				if lines[i][x] == byte('.') {
					continue
				}
				// fmt.Println("swapping at x=", x)	
				lines[i][j] = byte('O')
				lines[i][x] = byte('.')
				break
			}
		}
	}
	return lines
}

func ShiftZerosSouth(lines [][]byte) [][]byte {
	for i:= len(lines)-1 ; i >=0; i-- {
		for j := 0; j < len(lines[0]); j++ {
			if lines[i][j] != byte('.') {
				continue
			}
			for x:=i-1; x >= 0; x-- {
				if lines[x][j] == byte('#') {
					break
				}
				if lines[x][j] == byte('.') {
					continue
				}
				lines[i][j] = byte('O')
				lines[x][j] = byte('.')
				break
			}
		}
	}
	return lines
}

func ShiftZerosEast(lines [][]byte) [][]byte {
	for j:= len(lines[0])-1; j >0 ; j-- {
		for i :=0; i < len(lines); i++ {
			if lines[i][j] != byte('.'){
				continue
			}
			for x:=j-1; x>=0; x-- {
				if lines[i][x] == byte('#') {
					break
				}
				if lines[i][x] == byte('.') {
					continue
				}
				lines[i][j] = byte('O')
				lines[i][x] = byte('.')
				break
			}
		}
	}
	return lines
}

func CycleOnce(lines [][]byte) [][]byte {
	northShifted := ShiftZerosNorth(lines)
	westShifted :=ShiftZerosWest(northShifted)
	southShifted:=ShiftZerosSouth(westShifted)
	eastShifted :=ShiftZerosEast(southShifted)
	
	return eastShifted
}

func CalculateWeight(lines [][]byte) int {
	sum :=0
	for i:=0; i < len(lines); i++ {
		rocksInRow := 0
		for j:=0; j < len(lines[0]); j++ {
			if lines[i][j] == 'O' {
				rocksInRow+=1
			}
		}
		// fmt.Println("rocks in row", rocksInRow)
		// fmt.Println("dist", len(lines) -i)
		sum += (len(lines) - i) * rocksInRow
	}
	return sum
}

func Part1(lines [][]byte) int {
	// helper.PrintByteArray(lines)
	shifted := ShiftZerosNorth(lines)
	// fmt.Println("~~")
	// helper.PrintByteArray(shifted)
	weight := CalculateWeight(shifted)
	// fmt.Println(weight)
	return weight
}

func part2(lines [][]byte) int {
	line:=lines
	myMap := make(map[uint64]int)
	// NOTE: This was just a lucky mistake. Probably would need a hashmap of seen and distance to detect cycle, 
	for i :=0 ; i <1000; i++ {
		line = CycleOnce(line)
		// helper.PrintByteArray(line)
		// fmt.Println("~~")
		hash := helper.Hash2dByteArray(line)
		if myMap[hash] > 0 {
			fmt.Println("at", i,"hash: ", hash, "found at", myMap[hash])
		}
		myMap[hash] = i
	}
	return CalculateWeight(line)
}

func main() {
	stringLines := helper.ArrayFromFileLines()
	byteLines := helper.BytesFromStringArray(stringLines)
	// part1 := Part1(byteLines)
	// fmt.Println(part1)
	val:= part2(byteLines)
	fmt.Println(val)
}