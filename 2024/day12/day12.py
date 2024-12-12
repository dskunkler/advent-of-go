dir = [(0,1), (0,-1), (1,0), (-1,0)]

def calcPer(loc, crop, garden):
    row, col = loc
    sol = 0
    for dirRow, dirCol in dir:
        newRow,newCol = row +dirRow, col + dirCol
        if not 0<= newRow < len(garden) or not 0<= newCol < len(garden[0]) or garden[newRow][newCol] != crop:
            sol +=1
    return sol

def calcAreaAndPerimeter(crop, loc, garden, seen, currP, currA):
    row, col = loc
    print("loc", loc)
    if not 0 <= row < len(garden) or not 0 <= col < len(garden[0]) or loc in seen or garden[row][col] != crop:
        print("seen")
        return (0,0)
    seen.add(loc)
    currA += 1
    print("area", currA)
    currP += calcPer(loc, crop, garden)
    print("currP", currP)
    for rowDir, colDir in dir:
        newRow, newCol = row+rowDir, col+colDir
        newCurrA, newCurrP = calcAreaAndPerimeter(crop, (newRow, newCol), garden, seen, 0, 0)
        currP, currA= currP + newCurrP, currA + newCurrA
    print("returning", currA, currP)
    return (currA, currP)


def pt1():
    file = open("2024/day12/day12test.txt")
    garden = [[l for l in letter.strip()] for letter in file.readlines()]
    seen = set()
    print(garden)
    sol = 0
    for row in range(len(garden)):
        for col in range(len(garden[0])):
            area, perimeter = calcAreaAndPerimeter(garden[row][col], (row,col), garden, seen, 0, 0)
            print("perimeter and area for", garden[row][col], perimeter, area)
            sol += area * perimeter
            print("new sol", sol)

def isOOB(testLoc, crop, garden):
    testRow, testCol = testLoc
    return not 0 <= testRow < len(garden) or not 0<= testCol < len(garden[0]) or garden[testRow][testCol] != crop
def isInBound(loc, crop, garden):
    row,col = loc
    return 0<= row < len(garden) and 0 <= col < len(garden[0]) and garden[row][col]== crop

    
def calcTurns(loc, crop, garden):
    print("IN CALC TURNS")
    row, col = loc
    angles = 0
    left = (row, col-1)
    right = (row, col+1)
    up = (row-1, col)
    down = (row+1, col)
    upRight = (row-1, col+1)
    upLeft = (row-1, col-1)
    downRight= (row+1, col+1)
    downLeft = (row+1, col-1)

    # obtuse angles
    if isOOB(left, crop, garden) and isOOB(up, crop, garden):
        print("top left...")
        angles +=1
    if isOOB(left, crop, garden) and isOOB(down, crop, garden):
        print("bottom left...")
        angles+=1
    if isOOB(right, crop, garden) and isOOB(up, crop, garden):
        print("top right...")
        angles+=1
    if isOOB(right, crop, garden) and isOOB(down, crop, garden):
        print("bottom right...")
        angles+=1
    # acute angles
    if isInBound(up, crop, garden) and isInBound(right, crop, garden) and isOOB(upRight, crop, garden):
        angles+=1
    if isInBound(up, crop, garden) and isInBound(left, crop, garden) and isOOB(upLeft, crop, garden):
        angles+=1
    if isInBound(down, crop, garden) and isInBound(right, crop, garden) and isOOB(downRight, crop, garden):
        angles+=1
    if isInBound(down, crop, garden) and isInBound(left, crop, garden) and isOOB(downLeft, crop, garden):
        angles+=1
    
    
    
    print("angles from ", loc, angles)
    return angles

def calcAreaAndTurns(crop, loc, garden, seen, currP, currA):
    row, col = loc
    print("loc", loc)
    if not 0 <= row < len(garden) or not 0 <= col < len(garden[0]) or loc in seen or garden[row][col] != crop:
        print("seen")
        return (0,0)
    seen.add(loc)
    currA += 1
    print("area", currA)
    currP += calcTurns(loc, crop, garden)
    print("currP", currP)
    for rowDir, colDir in dir:
        newRow, newCol = row+rowDir, col+colDir
        newCurrA, newCurrP = calcAreaAndTurns(crop, (newRow, newCol), garden, seen, 0, 0)
        currP, currA= currP + newCurrP, currA + newCurrA
    print("returning", currA, currP)
    return (currA, currP)


def pt2():
    file = open("2024/day12/day12.txt")
    garden = [[l for l in letter.strip()] for letter in file.readlines()]
    seen = set()
    sol = 0
    for row in range(len(garden)):
        for col in range(len(garden[0])):
            area, perimeter = calcAreaAndTurns(garden[row][col], (row,col), garden, seen, 0, 0)
            print("angles and area for", garden[row][col], perimeter, area)
            sol += area * perimeter
            print("new sol", sol)

if __name__ == '__main__':
    pt2()