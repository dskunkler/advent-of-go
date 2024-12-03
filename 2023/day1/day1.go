package day1

import (
	"advent/helper"
	"fmt"
)


func Day1() {

    fileLines := helper.ArrayFromFileLines()

		total := 0

    for _, line := range fileLines {
        fmt.Println(line)
				temp := 0
				l:=0
				for l < len(line) {
					if line[l] >= '0' && line[l] <='9' {
						temp += 10 * int(line[l] - '0')
						break
					}
					l++
				}

				for r:=len(line)-1; r >=l; r-- {
					if line[r] > '0' && line[r] <= '9' {
						temp += int(line[r] - '0')
						// fmt.Println(temp)
						break
					}
				}

				total += temp
    }

		helper.Demo()

    fmt.Println(total)
}

func Part2() {
	fileLines := helper.ArrayFromFileLines()
	total := 0
	for _, line := range fileLines {
		i:=0
		// get first digit
		fmt.Println(line)
		for i < len(line) {
			d, error := helper.GetNumberFromRune(line[i])
			w, e := helper.GetNumberFromString(i,line)
			if error == nil {
				total+= 10*d
				fmt.Println("adding ", 10*d)
				break
			} else if e == nil {
				total+= 10*w
				fmt.Println("adding ", 10*w)
				break
			}
			i+=1
		}

		j:= len(line)-1
		for j >=i {
			d, error := helper.GetNumberFromRune(line[j])
			w, e := helper.GetNumberFromString(j,line)
			if error == nil {
				total+= d
				fmt.Println("adding ", d)
				break
			} else if e == nil {
				total+= w
				fmt.Println("adding ", w)
				break
			}
			j-=1
		}
	}
	fmt.Println(total)
}