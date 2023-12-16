package main

import (
	"advent/helper"
	"fmt"
)

type Key struct {
	Ipos int
	Jpos int
	Dir string
}

func handleMovement(myMap map[Key]int, elfMap []string, dir string, i, j int) map[Key]int {
	if i < 0 || j < 0 || i >= len(elfMap) || j >= len(elfMap[0]) || myMap[Key{i,j,dir}] > 0 {
		return myMap
	}
	myMap[Key{i,j,dir}] = 1
	posByte := elfMap[i][j]
	if posByte == byte('.') {
		if dir == "R" {
			myMap = handleMovement(myMap, elfMap, "R", i, j+1)
		} else if dir == "L" {
			myMap = handleMovement(myMap, elfMap, "L", i, j-1)
		} else if dir == "U" {
			myMap = handleMovement(myMap, elfMap, "U", i-1, j)
		} else if dir == "D" {
			myMap = handleMovement(myMap, elfMap, "D", i+1, j)
		}
	} else if posByte == byte('|') {
		if dir == "R"  || dir == "L" {
			myMap = handleMovement(myMap, elfMap, "U", i-1, j)
			myMap = handleMovement(myMap, elfMap, "D", i+1, j)
		} else if dir == "U" {
			myMap = handleMovement(myMap,elfMap , "U", i-1, j)
		} else {
			myMap = handleMovement(myMap, elfMap, "D", i+1, j)
		}
	} else if posByte == byte('-') {
		if dir == "U" || dir == "D" {
			myMap = handleMovement(myMap, elfMap, "L", i, j-1)
			myMap = handleMovement(myMap, elfMap,"R", i, j+1)
		} else if dir == "L" {
			myMap = handleMovement(myMap, elfMap, "L", i, j-1)
		} else {
			myMap = handleMovement(myMap, elfMap, "R", i, j+1)
		}
	} else if posByte == byte('\\') {
		if dir == "R" {
			myMap = handleMovement(myMap, elfMap, "D", i+1, j)
		} else if dir == "U" {
			myMap = handleMovement(myMap, elfMap, "L", i, j-1)
		} else if dir == "L" {
			myMap = handleMovement(myMap, elfMap, "U", i-1, j)
		} else {
			myMap = handleMovement(myMap, elfMap, "R", i, j+1)
		}
	} else if posByte == byte('/') {
		if dir == "R" {
			myMap = handleMovement(myMap, elfMap, "U", i-1, j)
		} else if dir == "D" {
			myMap = handleMovement(myMap, elfMap, "L", i, j-1)
		} else if dir == "L" {
			myMap = handleMovement(myMap, elfMap, "D", i+1, j)
		} else {
			myMap = handleMovement(myMap, elfMap, "R", i, j+1)
		}
	}

	return myMap
}

func handleShmovement( elfMap [][]byte, dir string, i, j int) [][]byte {
	if i < 0 || j < 0 || i >= len(elfMap) || j >= len(elfMap[0])  {
		return elfMap
	}
	posByte := elfMap[i][j]
	elfMap[i][j] = '#'
	if posByte == byte('.') {
		if dir == "R" {
			elfMap = handleShmovement( elfMap, "R", i, j+1)
		} else if dir == "L" {
			elfMap= handleShmovement( elfMap, "L", i, j-1)
		} else if dir == "U" {
			elfMap = handleShmovement( elfMap, "U", i-1, j)
		} else if dir == "D" {
			elfMap = handleShmovement( elfMap, "D", i+1, j)
		}
	} else if posByte == byte('|') {
		if dir == "R"  || dir == "L" {
			elfMap = handleShmovement( elfMap, "U", i-1, j)
			elfMap = handleShmovement( elfMap, "D", i+1, j)
		} else if dir == "U" {
			elfMap = handleShmovement(elfMap , "U", i-1, j)
		} else {
			elfMap = handleShmovement( elfMap, "D", i+1, j)
		}
	} else if posByte == byte('-') {
		if dir == "U" || dir == "D" {
			elfMap = handleShmovement( elfMap, "L", i, j-1)
			elfMap = handleShmovement( elfMap,"R", i, j+1)
		} else if dir == "L" {
			elfMap = handleShmovement( elfMap, "L", i, j-1)
		} else {
			elfMap = handleShmovement( elfMap, "R", i, j+1)
		}
	} else if posByte == byte('\\') {
		if dir == "R" {
			elfMap = handleShmovement( elfMap, "D", i+1, j)
		} else if dir == "U" {
			elfMap = handleShmovement( elfMap, "L", i, j-1)
		} else if dir == "L" {
			elfMap = handleShmovement( elfMap, "U", i-1, j)
		} else {
			elfMap = handleShmovement( elfMap, "R", i, j+1)
		}
	} else if posByte == byte('/') {
		if dir == "R" {
			elfMap = handleShmovement( elfMap, "U", i-1, j)
		} else if dir == "D" {
			elfMap = handleShmovement( elfMap, "L", i, j-1)
		} else if dir == "L" {
			elfMap = handleShmovement( elfMap, "D", i+1, j)
		} else {
			elfMap = handleShmovement( elfMap, "R", i, j+1)
		}
	}

	return elfMap
}

func getCount( myMap map[Key]int) int {
	newMap := map[[2]int]int {}
	for key, _ := range myMap {
		newMap[[2]int{key.Ipos, key.Jpos}] = 1
	}
	return len(newMap)
}
func part1()int {
	stringsArray := helper.ArrayFromFileLines()
	// byteLines := helper.BytesFromStringArray(stringsArray)
	myMap := map[Key]int{}
	myMap = handleMovement( myMap, stringsArray, "R", 0,0)
	ct:=  getCount(myMap)
	// helper.PrintByteArray(byteLines)
	return ct
}

func part2()int {
	sol :=0
	stringsArray := helper.ArrayFromFileLines()
	for i := 0; i < len(stringsArray) ; i ++ {
		myMap:= map[Key]int {}
		myMap2 := map[Key]int{}
		myMap = handleMovement(myMap, stringsArray, "R", i, 0)
		myMap2 = handleMovement(myMap2, stringsArray, "L", i, len(stringsArray[0])-1)
		ct :=getCount(myMap)
		sol = max(ct,sol)
		ct = getCount(myMap2)
		sol = max(ct,sol)
	}
	for j:= 0; j < len(stringsArray[0]); j++ {
		myMap:= map[Key]int {}
		myMap2 := map[Key]int{}
		myMap = handleMovement(myMap, stringsArray, "D", 0, j)
		myMap2 = handleMovement(myMap2, stringsArray, "U", len(stringsArray)-1, j)
		ct :=getCount(myMap)
		sol = max(ct,sol)
		ct = getCount(myMap2)
		sol = max(ct,sol)

	}
	return sol
}

func main() {
	// p1 := part1()
	// fmt.Println(p1)
	p2:= part2()
	fmt.Println(p2)
}