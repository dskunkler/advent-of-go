package helper

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


func Demo() {
	fmt.Println("good job")
}

func ArrayFromFileLines() []string {
	filePath := os.Args[1]
    readFile, err := os.Open(filePath)

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string

    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }

    readFile.Close()
		return fileLines
	}

func Max[T cmp.Ordered](x, y T) T {
  if x > y {
      return x
  }
  return y
} 


func GetNumberFromRune (r byte) (int, error) {
	if r >= '0' && r <='9' {
		return int(r - '0'), nil
	}
	return 0, fmt.Errorf("'%c' is not a number between 0 and 9", r)
}

func GetNumberFromString (i int, s string) (int, error) {
	words := map[string]int{
		"one" : 1,
		"two":2,
		"three":3,
		"four":4,
		"five":5,
		"six":6,
		"seven":7,
		"eight":8,
		"nine":9,
	}
	for key, val := range words{
		if strings.HasPrefix(s[i:], key) {
			return val, nil
		}
	}	
	return 0, fmt.Errorf("%s doesn't contain a number word starting at %d", s, i)
}

func IsNotNumberOrDot(r byte) bool {
	return (!(r >= '0' && r <='9') && r!='.')
}

func IsLeftRightOfLine(line string, left int, right int )bool{
	if (left > 0 && IsNotNumberOrDot(line[left-1])) || (right < len(line)-1 &&  IsNotNumberOrDot(line[right])) {			return true}
	return false
}

func NumberArrayFromString(line string)([]int, error) {
	digitRe:= regexp.MustCompile(`(-)?(\d+)`)
	strDigits := digitRe.FindAllString(line,-1)
	digits := make([]int,0)
	for _,s :=range strDigits {
		n,e :=strconv.Atoi(s)
		if e != nil {
			return digits,fmt.Errorf("Tried to convert not digit")
		}
		digits = append(digits,n )
	}
	return digits, nil
}

func GCD (a,b int)int {
	for a!=b {
		if a > b {
			a -=b
		} else {
			b-=a
		}
	}
	return a
}

func BytesFromStringArray(textLines []string) [][]byte {
	newArray := make([][]byte,0)
	for _,line := range textLines {
		tempLine:= make([]byte,0)
		for i:=0; i < len(line); i++ {
			tempLine = append(tempLine, line[i])
		}
		newArray = append(newArray, tempLine)
	}
	return newArray
}

func PrintByteArray ( lines [][]byte ) {
	for _, line := range lines {
		for _, b:= range line {
			fmt.Print(string(b), " ")
		}
		fmt.Println()
	}
}