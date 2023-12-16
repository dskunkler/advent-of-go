package main

import (
	"advent/helper"
	"fmt"
	"strconv"
	"strings"
)

func part2() int {
	stringLines := helper.ArrayFromFileLines()
	strArray := strings.Split(stringLines[0], ",")
	byteLines := helper.BytesFromStringArray(strArray)
	total:=0

	boxes := [256][]map[string]int{}

	for _, line := range byteLines {
		tempStr := string(line)
		if strings.Contains(tempStr, "=") {
			// do = stuff
			boxes = HandleEqualSign(boxes, tempStr)
		} else {
			// do - stuff
			boxes = HandleMinusSign(boxes, tempStr)
		}
	}

	for b, box := range boxes {
		for s, slot := range box {
			for _, val := range slot {
				currVal := (1+b) * (1+s) * val
				// fmt.Println("~~curr", currVal)
				total+= currVal
			}
		}
	}

	return total
}

func HandleMinusSign( boxes [256][]map[string]int, line string) [256][]map[string]int {
	part := line[:len(line)-1]
	hashedVal := HashFunc(0, []byte(part))
	contents := boxes[hashedVal]
	for i, m:= range contents {
		if m[part] > 0 {
			contents = append(contents[:i], contents[i+1:]...)
			boxes[hashedVal] = contents
			return boxes
		}
	}
	return boxes
}

func HandleEqualSign(boxes [256][]map[string]int, line string) [256][]map[string]int {
	parts := strings.Split(line, "=")
	hashedVal := HashFunc(0, []byte(parts[0]))
	contents := boxes[hashedVal]
	num, err := strconv.Atoi(parts[1])
	if err !=nil {
		fmt.Println("~~~~", err)
		return boxes
	}
	for i, m := range contents {
		// Already contains filter
		if m[parts[0]] > 0 {
				m[parts[0]] = num
				boxes[hashedVal][i] = m
				return boxes
		}
	}
	// Didnt find value
	boxes[hashedVal] = append(contents, map[string]int {parts[0]: num })
	return boxes
}

func part1()int  {
	stringLines := helper.ArrayFromFileLines()
	strArray := strings.Split(stringLines[0], ",")
	byteLines := helper.BytesFromStringArray(strArray)
	total := 0
	for _, line := range byteLines {
		curr := HashFunc(0, line)
		// fmt.Println("got", curr)
		total+= curr
	}
	return total
}

func HashFunc(currVal int, str []byte) int {
	for _, b := range str {
		currVal+=int(b)
		currVal*=17
		currVal%=256
	}
	return currVal
}
func main() {
	// res1:= part1()
	// fmt.Println("res1=", res1)
	 res2:=part2()
	 fmt.Println(res2)
}