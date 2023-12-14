package main

import (
	"advent/helper"
	"fmt"
)

func CheckVerticalSymmetry(lines []string,left, right int) (int,error) {
	// fmt.Println("right", right)
	// fmt.Println("left", left)
	for i:=0; i < len(lines); i++ {
		// fmt.Println("line", lines[i])
		for j := 0; (j + right < len(lines[0])) && left - j >= 0; j++ {
			if lines[i][right + j] != lines[i][left-j] {
				return 0, fmt.Errorf("Not veritcally symmetrical")
			}
		}
	}
	// fmt.Println("vertically symmetrical at ", right)
	return right, nil
}

func CheckHorizontalSymmetry(lines []string, top,bottom int) (int, error) {
	for j:=0; j < len(lines[0]); j++ {
		for i:=0; (bottom +i <len(lines)) && (top - i >=0); i++ {
			for lines[bottom+i][j] != lines[top - i][j] {
				return 0, fmt.Errorf("Not horizontally symmetrical")
			}
		} 
	}
	// fmt.Println("horizontally symmetrical at ", bottom)
	return bottom, nil
}


func SummarizePattern(linesArray [][]string, horizontalPrice, verticalPrice int) int {
	sum := 0
	for _, lines := range linesArray {
		// fmt.Println("getting summary for", lines)
		for i:= 1; i< len(lines); i++ {
			horizontalPos, err := CheckHorizontalSymmetry(lines,i-1, i )
			if err == nil {
				sum+= horizontalPos  * horizontalPrice
			}else {
				// fmt.Println(err)
			}
		}
		for j := 1; j < len(lines[0]); j++{
			verticalPos , err := CheckVerticalSymmetry(lines, j-1, j)
			if err == nil {
				sum += verticalPos * verticalPrice
			} else {
				// fmt.Println(err)
			}
		}
	}
	return sum
}

func ReverseByte(lines []string, a,b byte, i,j int) (byte, error){
	if  lines[i][j] == a {
		return b,nil
	}
	if lines[i][j] == b {
		return a,nil
	}
	return lines[i][j], fmt.Errorf(`wasnt {%w} or {%w} got {%w} instead`, a,b,lines[i][j])
}

// Need to finish this one still
// func SummarizeSmudge(linesArray[][]string, horizontalPrice, verticalPrice int) int {
// 	sum :=0
// 	for _, lines := range linesArray {
// 		copy := helper.CopyStringArray(lines)
// 		for i:=0; i < len(lines); i++ {
// 			for j:=0; j < len(lines[0]); j++ {
// 				reversed,err := ReverseByte(copy,'.', '#', i,j)
// 				if err == nil {
// 					bytes := []byte(lines[i])
// 					bytes[j] = reversed
// 					lines[i] = string(bytes)
// 					lines = copy
// 					val := SummarizePattern(lines, horizontalPrice,verticalPrice)
// 				}
// 			}
// 		}
// 	}
// }

func main() {
	textLines := helper.SplitFileByEmptyLine()
	res :=SummarizePattern(textLines, 100,1)
	fmt.Println("res=", res)
}