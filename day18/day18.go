package main

import (
	"advent/helper"
	"fmt"
	"strconv"
	"strings"
)

// make a template

func MakeImage(stringLines []string)[]string{
	r:=1
	d:=1
	for _, line := range stringLines {
		dir := string(line[0])
		temps := strings.Split(line, " ")
		fmt.Println(temps, len(temps))
		val,_ := strconv.Atoi(temps[1])
		if dir == "D" {
			d+=		val}
		if dir == "R" {
			r +=val
		}
	}
	image := make([]string,0)

	for i := 0; i < d; i++ {
		temp := make([]byte,0)
		for j:=0; j < r; j++ {
			temp = append(temp, '.')
		}
		image = append(image, string(temp))
	}
	return image
}

func travelImage(image []string, input []string) {
	image[0] = "#" + image[0][1:]
}

func part1() {
	stringLines := helper.ArrayFromFileLines()
	image :=MakeImage(stringLines)
	fmt.Println(image)
}

func main() {
	part1()
}