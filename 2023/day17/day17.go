package main

import (
	"advent/helper"
	"container/heap"
	"fmt"
	"strconv"
)

type Pos struct {
	Ipos,Jpos,Steps int
	Dir string
}

type IJ struct {
	Ipos, Jpos int
}
type Item struct {
	value Pos // position of last visited 
	priority int // priority in queue
	index int // index in heap
}

// PQ implementation from docs https://pkg.go.dev/container/heap#example-package-PriorityQueue
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	 return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Always want pop to give us the least heat loss so we use LT
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq PriorityQueue) Print() {
	for _,item := range pq {
		fmt.Print(item.priority, ": ", item.value, ", ")
	}
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value Pos, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}



func pq1() int {
	stringLines := helper.ArrayFromFileLines()
	pq := make(PriorityQueue,0)
	seen := make(map[Pos]int)
	rightNum, _ := strconv.Atoi(string(stringLines[0][1]))
	downNum, _ := strconv.Atoi(string(stringLines[1][0]))
	i:=0
	startRight := &Item{
		priority: rightNum,
		value: Pos{0,1,1,"R"},
	}
	startDown := &Item{
		priority: downNum,
		value: Pos{1,0,1,"D"},
	}
	pq.Push(startRight)
	pq.Push(startDown)
	heap.Init(&pq)
	
	// Get the least thermal loss item
	for pq.Len() > 0 {
		// if i >20 {
		// 	os.Exit(1)
		// }
		i+=1
		// pq.Print()
		// Get the lowest value off the queue
		item:= heap.Pop(&pq).(*Item)

			// Set the index as expored
		// ij :=IJ{item.value.Ipos, item.value.Jpos}
		if seen[item.value] > 0 { 
			continue
		}
		seen[item.value] = 1


		i,j,stepsTaken,dir,p:= item.value.Ipos, item.value.Jpos, item.value.Steps,item.value.Dir, item.priority
		// If we got the solution, return it
		// fmt.Println("Current pos", i, j, stepsTaken,dir)
		// fmt.Println("Curr priority", item.priority)
		if i == len(stringLines)-1 && j == len(stringLines[0])-1 && stepsTaken>=4 {
			fmt.Println("Adding final", p)
			return p
		}

		if dir == "R" {
			nextPos:= Pos{i, j+1, stepsTaken+1, "R"}
			if stepsTaken < 10 && j+1 < len(stringLines[0]) && seen[nextPos] < 1{
				currVal, _ := strconv.Atoi(string(stringLines[i][j+1]))
				// fmt.Println("Adding", currVal)
				newPriority := item.priority + currVal
				newRight := &Item {
					priority: newPriority,
					value: nextPos,
				}
				pq.Push(newRight)
			}
			if stepsTaken >= 4 {

				nextPos =  Pos{i-1, j, 1, "U"}
				if i-1 >=0  && seen[nextPos] <1 {
					currVal, _ := strconv.Atoi(string(stringLines[i-1][j]))
					// fmt.Println("Adding", currVal)
					newPriority := item.priority + currVal
					newUp := &Item {
							priority: newPriority,
							value:nextPos,
						}
					pq.Push(newUp)
				}
				nextPos =  Pos{i+1, j, 1, "D"}
				if i+1 < len(stringLines) && seen[nextPos] < 1 {
					currVal, _ := strconv.Atoi(string(stringLines[i+1][j]))
					// fmt.Println("Adding", currVal)
					newPriority := item.priority + currVal
	
					newDown := &Item {
						priority: newPriority,
						value:nextPos,
					}
					pq.Push(newDown)
				}
			}
		}
		if dir == "D" {
			// fmt.Println("Going D")
			nextPos := Pos{i+1, j, stepsTaken+1, "D"}
			if stepsTaken < 10 && i+1 < len(stringLines) && seen[nextPos] < 1 {	
				currVal, _ := strconv.Atoi(string(stringLines[i+1][j]))
				// fmt.Println("Adding", currVal)
				newPriority := item.priority + currVal

				newDown := &Item {
					priority: newPriority,
					value: nextPos,
				}
				pq.Push(newDown)
			} 
			if stepsTaken >= 4 {
				nextPos = Pos{i, j-1,1,"L"}
				// fmt.Println("Seen? j-1? ", seen[nextPos])
				if j-1 >=0 && seen[nextPos] < 1 {
					// fmt.Println("Adding L")
					currVal, _ := strconv.Atoi(string(stringLines[i][j-1]))
					// fmt.Println("Adding", currVal)
					newPriority := item.priority + currVal
	
					newLeft := &Item {
						priority: newPriority,
						value: nextPos,
					}
					pq.Push(newLeft)
				} 
				nextPos = Pos{i, j+1, 1, "R"}
				// fmt.Println("Seen j+1?", seen[nextPos])
				if j+1 < len(stringLines[0]) && seen[nextPos] < 1 {
					// fmt.Println("adding R")
					currVal, _ := strconv.Atoi(string(stringLines[i][j+1]))
					// fmt.Println("Adding", currVal)
					newPriority := item.priority + currVal
					newRight := &Item {
						priority: newPriority,
						value: nextPos,
					}
					pq.Push(newRight)
				}
			}
		}
		if dir == "L" {
			// fmt.Println("Going L")
			nextPos := Pos{i, j-1,stepsTaken+1,"L"}
			if stepsTaken < 10 && j-1 >=0 && seen[nextPos] < 1 {
				currVal, _ := strconv.Atoi(string(stringLines[i][j-1]))
				// fmt.Println("Adding", currVal)
				newPriority := item.priority + currVal
			
				newLeft := &Item {
					priority: newPriority,
					value: nextPos,
				}
				pq.Push(newLeft)
			}
			if stepsTaken >=4 {
				nextPos = Pos{i-1, j, 1, "U"}
				if i-1 >=0 && seen[nextPos] < 1 {
					currVal, _ := strconv.Atoi(string(stringLines[i-1][j]))
					// fmt.Println("Adding", currVal)
					newPriority := item.priority + currVal
	
					newUp := &Item {
						priority: newPriority,
						value: nextPos,
					}
					pq.Push(newUp)
				}
				nextPos = Pos{i+1, j, 1, "D"}
				if i+1 < len(stringLines) && seen[nextPos] < 1 {
					currVal, _ := strconv.Atoi(string(stringLines[i+1][j]))
					// fmt.Println("Adding", currVal)
					newPriority := item.priority + currVal
	
					newDown := &Item {
						priority: newPriority,
						value: nextPos,
					}
					pq.Push(newDown)
				}
			}
			
		}
		if dir == "U" {
			// fmt.Println("Going U")
			nextPos := Pos{i-1, j, stepsTaken+1, "U"}
			if stepsTaken < 10 && i-1 >= 0 && seen[nextPos] < 1 {
				currVal, _ := strconv.Atoi(string(stringLines[i-1][j]))
				// fmt.Println("Adding", currVal)
				newPriority := item.priority + currVal

				newUp := &Item {
				priority: newPriority,
				value: nextPos,
				}
				pq.Push(newUp)
			}
			if stepsTaken >=4 {

				nextPos = Pos{i, j-1,1,"L"}
				if j-1 >=0 && seen[nextPos] < 1 {
					currVal, _ := strconv.Atoi(string(stringLines[i][j-1]))
					// fmt.Println("Adding", currVal)
					newPriority := item.priority + currVal
	
					newLeft := &Item {
					priority: newPriority,
					value: nextPos,
					}
					pq.Push(newLeft)
				}
				nextPos = Pos{i, j+1, 1, "R"}
				if j+1 < len(stringLines[0]) && seen[nextPos] < 1 {
					currVal, _ := strconv.Atoi(string(stringLines[i][j+1]))
					// fmt.Println("Adding", currVal)
					newPriority := item.priority + currVal
	
					newRight := &Item {
						priority: newPriority,
						value: nextPos,
					}
					pq.Push(newRight)
				}
			}
		}

	heap.Init(&pq)
	}
	return -1
}

// func djikstras() {
// 	// enqueue the first two and mark 00 as visited
// 	// for each node  in the queuquq
// 	// pop
// 	// if visited (i,j,dir) continue
// 	// mark as visited (i, j, stepsTaken, dir)
// 	// pay the price to move, and add appropriate directions


// }

func main() {
	part1 := pq1()
	fmt.Println(part1)
}