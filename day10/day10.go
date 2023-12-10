package day10

import (
	"advent/helper"
	"fmt"
)

var directions =map[byte][2][2]int {
	'|': [2][2]int {{-1,0},{1,0}},
	'-': [2][2]int {{0,-1},{0,1}},
	'L': [2][2]int {{-1,0}, {0,1}},
	'J': [2][2]int {{-1,0}, {0,-1}},
	'7': [2][2]int {{1,0}, {0,-1}},
	'F': [2][2]int {{1,0}, {0,1}},
}

func Part1() {
	textLines := helper.ArrayFromFileLines()
	byteLines := helper.BytesFromStringArray(textLines)
	fmt.Println(textLines)
	sol := 0
	startPos:= make([]int,0)
	// find start
	for i := 0; i < len(textLines); i++ {
		for j :=0; j < len(textLines[i]); j++ {
			if textLines[i][j] == 'S' {
				startPos = append(startPos,i )
				startPos = append(startPos, j)
				break
			}
		}
	}
	fmt.Println("startpos", startPos)
	// Add directions to queue
	queue:= make([][2]int,0)
	if startPos[0] > 0 {
		aboveByte := byteLines[startPos[0]-1][startPos[1]]
		// Above character must be |, 7 , or F go south
		if aboveByte == '|' || aboveByte == '7' || aboveByte == 'F'{
			up:= [2]int {
				startPos[0]-1, startPos[1],
			}
			queue = append(queue,up)
		}
	}
	if startPos[1] < len(byteLines[0]) -1 {
		// right char must go to west
		rightByte := byteLines[startPos[0]][startPos[1]+1]
		if rightByte == '-' || rightByte=='7' || rightByte == 'J'{
			right:=[2]int {
				startPos[0], startPos[1]+1,
			}
			queue = append(queue, right)
		}
	 }

	 if startPos[0] < len(byteLines)-1 {
		// down char must go north
		downByte :=byteLines[startPos[0]+1][startPos[1]]
		if downByte == '|' || downByte == 'J' || downByte == 'L' {
			down:=[2]int {
				startPos[0]+1, startPos[1],
			}
			queue = append(queue, down)
		}
	 }

	 if startPos[1] > 0 {
		// left char must go east
		leftByte := byteLines[startPos[0]][startPos[1]-1]
		if leftByte == '-' || leftByte == 'L' || leftByte == 'F' {
			left:=[2]int{
				startPos[0],startPos[1]-1,
			}
			queue = append(queue, left)
		}
	 }
	 fmt.Println("start q", queue)
	// bfs from start
	byteLines[startPos[0]][startPos[1]] = '*'

	depth := 0
	// we dont ever want to see the same spot unless we're hitting the middle of the circle
	// therefore, clear non pipes in start
	for i:= len(queue)-1; i >=0; i-- {
		q := queue[i]
		if byteLines[q[0]][q[1]] == '*' || byteLines[q[0]][q[1]] == '.' {
			queue = append(queue[:i],queue[i+1:]... )
		}
	}
	// fmt.Println("starting queue",queue)

	// while we have a queue
	for len(queue) > 0 {
		// add depth
		depth +=1
		// iterate over the current queue
		tempQ := queue
		queue = make([][2]int,0)
		// fmt.Println("depth",depth)
		// while we have a temp q
		for len(tempQ) > 0 {
			node := tempQ[0]
			// seperate the node from the temp queue
			tempQ = tempQ[1:]

			// get current pos
			x,y := node[0],node[1]

			// get the byte
			tempByte := byteLines[x][y]
			if tempByte == '*' || tempByte == '.' {
				continue
			}
			
			// mark the position as seen
			byteLines[x][y] = '*'
			// fmt.Println(string(tempByte))

			// Find the directions to go 
			dir1 := directions[tempByte]
			// fmt.Println("dir1", dir1)

			// If the directions are both seen,  exit
			if byteLines[x+dir1[0][0]][y +dir1[0][1]] == '*' && byteLines[x + dir1[1][0]][y +dir1[1][1]] == '*' {
				fmt.Println("found from both sides")
				sol = depth
				queue = make([][2]int, 0)
				break
			}
			
			// get the new positions
			newPos1 := [2]int {
				x + dir1[0][0], y+dir1[0][1],
			}
			newPos2 := [2]int {
				x+dir1[1][0], y+dir1[1][1],
			}

			// Add them to the queue
			queue = append(queue, newPos1)
			queue = append(queue, newPos2)
		}
		// helper.PrintByteArray(byteLines)
	}

	fmt.Println(sol)
}