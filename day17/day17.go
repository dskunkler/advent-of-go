package main

import (
	"advent/helper"
	"fmt"
	"math"
	"os"
	"strconv"
)

type IJKey struct {
	IPos int
	JPos int
}

// Returns heat Total
func dfs(explored map[IJKey]int, input []string, dir string,bestSeen *float64, currLoss float64, i, j, stepsTaken int) float64 {
	if i < 0 || j < 0 || i >= len(input) || j >= len(input[0]) || explored[IJKey{i, j}] == 1 || currLoss > *bestSeen || input[i][j] == '*' {
		// fmt.Println("Already seen at", i, j)
		return math.MaxFloat64
	}
	// fmt.Println("visiting", i,j)
	// fmt.Println("ex", explored)
	// fmt.Println("best seen", *bestSeen)
	res := math.MaxFloat64
	currVal,_ := strconv.Atoi(string(input[i][j]))
	// fmt.Println("adding", currVal)
	currLoss += float64(currVal)
	if i == len(input)-1 && j == len(input[0])-1 {
		fmt.Println("FOUND IT", currLoss)
		helper.PrintStringArray(input)
		return currLoss
	}
	inputCopy := helper.CopyStringArray(input)
	inputCopy[i] = inputCopy[i][:j] + "*" + inputCopy[i][j+1:]
	explored[IJKey{i, j}] = 1
	if dir == "R" {
		// fmt.Println("R, steps", stepsTaken, "Pos ", i, j)
		if stepsTaken < 3 {
			tempStraight := dfs(explored, inputCopy, "R",bestSeen,currLoss, i, j+1, stepsTaken+1)
			res = math.Min(res, tempStraight)
			*bestSeen = math.Min(res,*bestSeen)
		}
		tempUp := dfs(explored, inputCopy, "U",bestSeen,currLoss, i-1, j, 1)
		res = math.Min(res, tempUp)
		*bestSeen = math.Min(res,*bestSeen)
		tempDown := dfs(explored, inputCopy, "D",bestSeen,currLoss, i+1, j, 1)
		res = math.Min(res, tempDown)
		*bestSeen = math.Min(res,*bestSeen)

	}
	if dir == "D" {
		// fmt.Println("D,steps ", stepsTaken, "Pos ",  i , j)
		if stepsTaken < 3 {
			tempStraight := dfs(explored, inputCopy, "D",bestSeen, currLoss, i+1, j, stepsTaken+1)
			res = math.Min(res, tempStraight)
		}
		tempLeft := dfs(explored, inputCopy, "L",bestSeen, currLoss,i, j-1,  1)
		res = math.Min(res, tempLeft)
		*bestSeen = math.Min(res,*bestSeen)

		tempRight := dfs(explored, inputCopy, "R",bestSeen, currLoss,i, j+1,  1)
		res = math.Min(res, tempRight)
	}
	if dir == "L" {
		// fmt.Println("L, steps", stepsTaken, "Pos, ", i,j)
		if stepsTaken < 3 {
			tempStraight := dfs(explored, inputCopy, "L",bestSeen, currLoss,i, j-1,  stepsTaken+1)
			res = math.Min(res, tempStraight)
			*bestSeen = math.Min(res,*bestSeen)

		}
		tempUp := dfs(explored, inputCopy, "U",bestSeen, currLoss,i-1, j,  1)
		res = math.Min(res, tempUp)
		*bestSeen = math.Min(res,*bestSeen)
		tempDown := dfs(explored, inputCopy, "D",bestSeen, currLoss,i+1, j,  1)
		res = math.Min(res, tempDown)
	}
	if dir == "U" {
		// fmt.Println("U, steps", stepsTaken, "pos, ", i,j)
		if stepsTaken < 3 {
			tempStraight := dfs(explored, inputCopy, "U",bestSeen, currLoss,i-1, j,  stepsTaken+1)
			res = math.Min(res, tempStraight)
			*bestSeen = math.Min(res,*bestSeen)

		}
		tempLeft := dfs(explored, inputCopy, "L",bestSeen, currLoss,i, j-1,  1)
		res = math.Min(res, tempLeft)
		*bestSeen = math.Min(res,*bestSeen)

		tempRight := dfs(explored, inputCopy, "R",bestSeen, currLoss,i, j+1,  1)
		res = math.Min(res, tempRight)
	}
	explored[IJKey{i, j}] = 0
	return res
}

func part1() float64 {
	// dfs, keep track of visited positions,
	explored := map[IJKey]int{}
	input := helper.ArrayFromFileLines()
	// fmt.Println(input)
	// fmt.Println(input[0][0])
	bestSeen :=math.MaxFloat64
	startRight  := dfs(explored, input, "R",&bestSeen, 0.0,0,1,1 )
	fmt.Println("Found first answer...")
	os.Exit(1)
	startDown := dfs(explored,input, "D",&bestSeen,0.0, 1,0,1 )
	return math.Max(float64(startDown),float64(startRight))
}

func main() {
	fmt.Println(part1())
}