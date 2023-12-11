package day10

import (
	"advent/helper"
	"fmt"
)

var directions =map[byte][2][2]int {
	'|':  {{-1,0},{1,0}},
	'-':  {{0,-1},{0,1}},
	'L':  {{-1,0}, {0,1}},
	'J':  {{-1,0}, {0,-1}},
	'7':  {{1,0}, {0,-1}},
	'F':  {{1,0}, {0,1}},
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
	}
	helper.PrintByteArray(byteLines)

	fmt.Println(sol)
}

func markXforPos(lines[][]byte,lineCopy[][]byte,currX,currY, x,y int) ([][]byte,int) {
	// always want to check right
				// if _, we need to check if were coming from the left, then we'll look below
			// if _ and we're coming from the right, we'll look above 
			tiles:=0
			fmt.Println(string(lineCopy[x][y]))
			if lineCopy[x][y] == '-' {
				if y > currY {
				// we're coming from left, looking below
					if x+1 < len(lines) {
						if lines[x+1][y] != '*'{
							tiles+=1
							lines[x+1][y] ='X'
						}
					}
				} else {
					// were coming from the right, looking above
					if x -1 >=0 {
						if lines[x-1][y] == '.'{
							tiles+=1
							lines[x-1][y] = 'X'
						}
					}
				}
			}
			// if |, we need to check right if going bottem to top
			// or left if going top to bottom
			if lineCopy[x][y] == '|'{
				if x > currX {
					// top to bottom, check left
					if y - 1 >=0 && lines[x][y-1] != '*' {
						tiles+=1
							lines[x][y-1] = 'X'
						}
				} else {
					if y + 1 < len(lines[x])  && lines[x][y+1] != '*'{
						tiles+=1
						lines[x][y+1] = 'X'
					}
				}
			}
			// if F we need to check obtuse (acute always looks at pipe)
			if lineCopy[x][y] == 'F' {
				fmt.Println("F")
				if y > currY {
					// were moving right to left
					fmt.Println("right to left")
					// Check top
					if x - 1 >= 0  && lines[x-1][y] != '*' {
						fmt.Println("foudn above")
						tiles+=1
						lines[x-1][y] = 'X'
					}
					// Check left
					if y -1 >=0 && lines[x][y-1] != '*' {
						fmt.Println("found left")
						tiles+=1
						lines[x][y-1]= 'X'
					}
				}
			}
			// if L we want to check right if coming from top
			if lineCopy[x][y] == 'L'{
				fmt.Println("L")
				if x < currX {
					if y - 1 >=0 && lines[x][y-1] != '*'{
						// fmt.Println()
						lines[x][y-1] = 'X'
						tiles+=1
						
					}
					if x + 1 < len(lines) && lines[x+1][y] != '*'{
						tiles+=1
						lines[x+1][y] = 'X'
					}
				}
			}
			if lineCopy[x][y] == 'J' {
				if currY < y {
					// coming from the left
					if x+1 < len(lines) && lines[x+1][y] != '*'{
						tiles+=1
						lines[x+1][y ] = 'X'
					}
					if y+1 < len(lines[0]) && lines[x][y+1] != '*' {
						tiles+=1
						lines[x][y+1] = 'X'
					}
				}
			}
			if lineCopy[x][y] == '7' {
				if currX > x {
					// coming from bottom
					if y+1 < len(lines[x]) && lines[x][y+1] != '*' {
						tiles+=1
						lines[x][y+1]='X'
					}
					if x-1 > 0 && lines[x-1][y] != '*'{
						tiles+=1
						lines[x-1][y] = 'X'
					}
				}
			}
			return lines, tiles
}

func dfsTiles(lines [][]byte, x,y int) ([][]byte, int){
	tiles:=0

	if x < 0 || y < 0 || x >= len(lines) || y >= len(lines[0]) || lines[x][y] != '.' {
		fmt.Println(x, y, "= ", string(lines[x][y]))
		return lines,tiles
	}
	fmt.Println("~~adding a new X")
	tiles+=1
	lines[x][y] = 'X'
	var tiles1, tiles2, tiles3,tiles4 int
	lines,tiles1 = dfsTiles(lines, x+1,y)
	lines, tiles2 = dfsTiles(lines,x-1,y)
	lines, tiles3 = dfsTiles(lines,x,y+1)
	lines, tiles4 = dfsTiles(lines,x,y-1)
	tiles += tiles1 + tiles2 + tiles3 +tiles4
	// fmt.Println("returning ", tiles)
	return lines,tiles

}

func Part2() {
	textLines := helper.ArrayFromFileLines()
	byteLines := helper.BytesFromStringArray(textLines)
	bytesCopy := helper.CopyTwoDBytesArray(byteLines)
	// fmt.Println(textLines)
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
	// ADD STARTS
	// if startPos[1] < len(byteLines[0]) -1 {
	// 	// right char must go to west
	// 	rightByte := byteLines[startPos[0]][startPos[1]+1]
	// 	if rightByte == '-' || rightByte=='7' || rightByte == 'J'{
	// 		right:=[2]int {
	// 			startPos[0], startPos[1]+1,
	// 		}
	// 		queue = append(queue, right)
	// 	}
	//  }
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
	// I know my start so I'm just skipping those so I can go just one direction and do some special stuff for my dfs
	 fmt.Println("start q", queue)
	// bfs from start
	byteLines[startPos[0]][startPos[1]] = '*'

	depth := 0
	tiles:=0
	// while we have a queue
	for len(queue) > 0 {
		// add depth
		depth +=1
		// iterate over the current queue
		tempQ := queue
		queue = make([][2]int,0)
		// fmt.Println("depth",depth)
		// while we have a temp q

		// 
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
	}
	fmt.Println("byteCopy")
	helper.PrintByteArray(bytesCopy)
	// Do the queue again with the copy
	// instead of looking for a . just mark it if its NOT a * in the original
	fmt.Println("Startpos", startPos)
	 if startPos[0] < len(byteLines)-1 {
		// down char must go north
		downByte :=bytesCopy[startPos[0]+1][startPos[1]]
		fmt.Println("add to queue",string(downByte))
		if downByte == '|' || downByte == 'J' || downByte == 'L' {
			down:=[2]int {
				startPos[0]+1, startPos[1],
			}
			queue = append(queue, down)
		}
	 }
fmt.Println("new queue", queue)
	for len(queue) > 0 {
		// fmt.Println("running the new queue")
		// add depth
		depth +=1
		// iterate over the current queue
		tempQ := queue
		queue = make([][2]int,0)
		// fmt.Println("depth",depth)
		// while we have a temp q

		// 
		for len(tempQ) > 0 {
			node := tempQ[0]
			// seperate the node from the temp queue
			tempQ = tempQ[1:]

			// get current pos
			x,y := node[0],node[1]

			// get the byte
			tempByte := bytesCopy[x][y]
			if tempByte == '*' || tempByte == '.' {
				continue
			}
			
			// mark the position as seen
			bytesCopy[x][y] = '*'
			// fmt.Println(string(tempByte))

			// Find the directions to go 
			dir1 := directions[tempByte]
			// fmt.Println("dir1", dir1)

			// If the directions are both seen,  exit
			if bytesCopy[x+dir1[0][0]][y +dir1[0][1]] == '*' && bytesCopy[x + dir1[1][0]][y +dir1[1][1]] == '*' {
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

			byteLines, tiles1 := markXforPos(byteLines,bytesCopy, x,y,newPos1[0], newPos1[1])
	
			byteLines, tiles2 := markXforPos(byteLines, bytesCopy, x,y, newPos2[0], newPos2[1])
			// helper.PrintByteArray(byteLines)
			tiles+=tiles1
			tiles+=tiles2

			// Add them to the queue
			queue = append(queue, newPos1)
			queue = append(queue, newPos2)
		}
	}
	fmt.Println("tiles",tiles)
	helper.PrintByteArray(byteLines)
	var tempTiles int
	for i:= 0; i < len(byteLines); i++ {
		for j:=0; j < len(byteLines[0]); j++ {
			if byteLines[i][j] == 'X' {
				fmt.Println("found an X @", i,j)
				byteLines,tempTiles  = dfsTiles(byteLines,i-1,j)
				byteLines,tempTiles  = dfsTiles(byteLines,i+1,j)
				byteLines,tempTiles  = dfsTiles(byteLines,i,j-1)
				byteLines,tempTiles  = dfsTiles(byteLines,i,j+1)
				tiles+= tempTiles
				fmt.Println("foudn temp", tempTiles)
				fmt.Println("tiles", tiles)
			}
		}
	}
	fmt.Println("new tiles", tiles)
	helper.PrintByteArray(byteLines)


	fmt.Println(sol)
}