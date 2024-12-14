import re

def calcFinalPos(currX, currY, xVel, yVel, width, height, seconds):
    maxW, maxH = width-1, height-1
    # print("maxw ", maxW)
    # print("maxh ", maxH)
    for i in range(seconds):
        # print("curr ",(currX, currY))
        currX += xVel
        currY += yVel
        # print("new", (currX, currY))
        if currX < 0:
            # print("x less 0")
            currX = maxW + currX + 1
        if currX > maxW:
            # print("greater than max x")
            currX = currX - maxW -1
        if currY < 0:
            # print("y less 0")
            currY = maxH + currY + 1
            # print("new y", currY)
        if currY > maxH:
            # print("greater than max y")
            currY = currY - maxH -1
            # print('new y', currY)
        # print("wrapped ", (currX, currY))
    return (currX, currY)

def calcMultiple(quad):
    start = quad[0]
    for i in range(1, 4):
        start*= quad[i]
        print(start)
    return start

def numNeighbors(row, col, grid):
    return 0 if not 0<= row < len(grid) or not 0 <= col < len(grid[0]) else grid[row][col]
def calcNeighborScore(grid):
    neighbors = [(0,1), (1,1),(1,0),(1,-1),(-1, -1), (0, -1),(-1,0),  (-1,1) ]
    sol = 0
    for row in range(len(grid)):
        for col in range(len(grid[0])):
            if grid[row][col] !=0:
                for x,y in neighbors:
                    sol += numNeighbors(row+x, col+y, grid)
    return sol


def pt1():
    file = open("2024/day14/day14.txt")
    positions = [line.strip() for line in file.readlines()]
    # grid = [[0 for i in range(101)] for j in range(103)]
    # # print(grid)
    quad = [0 for i in range(4)]
    width, height , seconds= 101, 103, 1
    robots = []
    for line in positions:
        # print(line)
        split = re.split("\s", line)
        startingPos = re.findall("(\d+),(\d+)", split[0])[0]
        vel = re.findall("(-?\d+),(-?\d+)", split[1])[0]
        robots.append((startingPos, vel))
    sol = (0, 0)
    for i in range(100000):
        newGrid = [[0 for i in range(101)] for j in range(103)]
        newRobots = []
        for startingPos, vel in robots:
            loc = calcFinalPos(int(startingPos[0]), int(startingPos[1]), int(vel[0]),int(vel[1]),width, height, seconds)
            newRobots.append((loc, vel))
            # print(loc)
            newGrid[loc[1]][loc[0]] +=1
        weight = calcNeighborScore(newGrid)
        currMax = sol[0]
        sol = sol if currMax > weight else (weight, i)
        print("new sol", sol)
        robots = newRobots

    print("sol", sol)


            


    
def pt2():
    file = open("2024/day14/day14.txt")

if __name__ == '__main__':
    pt1()