def findGuardPos(roomMap):
    guard = {'^', '>', '<', 'V'}
    for row in range(len(roomMap)):
        for col in range(len(roomMap[0])):
            if roomMap[row][col] in guard:
                return (row, col)
def moveGuard(currPos, dir, roomMap, visited):
    # print("cur pos", currPos)
    rowDif, colDif = dir
    row,col = currPos
    while row >= 0 and row < len(roomMap) and col >= 0 and col < len(roomMap[0]):
        visited.add((row,col))
        # print(" cur = ",row,col)
        nextRow, nextCol = row + rowDif, col + colDif
        # print("next", nextRow, nextCol)
        if (nextRow >= 0 and nextRow < len(roomMap) and nextCol >= 0 and nextCol < len(roomMap[0])) and roomMap[nextRow][nextCol] == '#':
            # print("next is obs")
            # V
            if rowDif > 0:
                # print("go left")
                rowDif, colDif = 0,-1
                nextRow, nextCol = row + rowDif, col + colDif
                # ^
            elif rowDif < 0:
                # print("go right")
                rowDif, colDif = 0,1
                nextRow, nextCol = row + rowDif, col + colDif
                # <
            elif colDif < 0:
                # print("go up")
                rowDif, colDif = -1,0
                nextRow, nextCol = row + rowDif, col + colDif
                # >
            else:
                # print("go down")
                rowDif, colDif = 1,0
                nextRow, nextCol = row + rowDif, col + colDif
        row, col = nextRow, nextCol
def hasCycle(currPos, dir, roomMap, visited):
    # print("cur pos", currPos)
    rowDif, colDif = dir
    row,col = currPos
    while row >= 0 and row < len(roomMap) and col >= 0 and col < len(roomMap[0]):
        if (row,col, rowDif, colDif) in visited:
            # print("found cycle here ", (row,col, rowDif, colDif))
            return 1
        visited.add((row,col, rowDif, colDif))
        # print(" cur = ",row,col)
        nextRow, nextCol = row + rowDif, col + colDif
        # print("next", nextRow, nextCol)
        if (nextRow >= 0 and nextRow < len(roomMap) and nextCol >= 0 and nextCol < len(roomMap[0])) and roomMap[nextRow][nextCol] == '#':
            # print("next is obs")
            # V
            if rowDif > 0:
                # print("go left")
                rowDif, colDif = 0,-1
                # nextRow, nextCol = row + rowDif, col + colDif
                # ^
            elif rowDif < 0:
                # print("go right")
                rowDif, colDif = 0,1
                # nextRow, nextCol = row + rowDif, col + colDif
                # <
            elif colDif < 0:
                # print("go up")
                rowDif, colDif = -1,0
                # nextRow, nextCol = row + rowDif, col + colDif
                # >
            else:
                # print("go down")
                rowDif, colDif = 1,0
                # nextRow, nextCol = row + rowDif, col + colDif
            # print("next..", nextRow, nextCol)
        else:
            # print("go this way")
            row, col = nextRow, nextCol
    return 0
        
def placeObjects(pos, dir, roomMap, visited):
    # take my visited... See if theres any position in there that I can swap to a # to create a cycle
    sol = 0
    print("visited length without start", len(visited))
    for row, col in visited:
        if (row, col) == pos:
            continue
        tempVisited = set()
        # print("row, col", row, col)
        roomMap[row] = roomMap[row][:col]  + '#' + roomMap[row][col+1:]
        # print("roroommaprow", roomMap[row])
        sol += hasCycle(pos, dir, roomMap, tempVisited)
        # roomMap[row][col] = '.'
        roomMap[row] = roomMap[row][:col]  + '.' + roomMap[row][col+1:]
        # print("roroommaprow", roomMap[row])
    print("sol = ",sol)

def pt1():
    file = open("2024/day6/day6test.txt")
    roomMap = [line.strip() for line in file.readlines()]
    visited = set()
    initialPos = findGuardPos(roomMap)
    guardDir =roomMap[initialPos[0]][initialPos[1]] 
    if guardDir == '^':
        moveGuard(initialPos, (-1,0), roomMap, visited)
    elif guardDir == 'V':
        moveGuard(initialPos, (1,0), roomMap, visited)
    if guardDir == '<':
        moveGuard(initialPos, (0,-1), roomMap, visited)
    if guardDir == '>':
        moveGuard(initialPos, (0,1), roomMap, visited)
    print(visited)
    print(len(visited))  
                    
def pt2():
    file = open("2024/day6/day6.txt")
    roomMap = [line.strip() for line in file.readlines()]
    # line = file.readline()
    visited = set()
    initialPos = findGuardPos(roomMap)
    guardDir =roomMap[initialPos[0]][initialPos[1]] 
    if guardDir == '^':
        moveGuard(initialPos, (-1,0), roomMap, visited)
        # print(visited)
        placeObjects(initialPos, (-1,0), roomMap, visited)
    elif guardDir == 'V':
        moveGuard(initialPos, (1,0), roomMap, visited)
        # print(visited)
        placeObjects(initialPos, (1,0), roomMap, visited)
    if guardDir == '<':
        moveGuard(initialPos, (0,-1), roomMap, visited)
        # print(visited)
        placeObjects(initialPos, (0,-1), roomMap, visited)
    if guardDir == '>':
        moveGuard(initialPos, (0,1), roomMap, visited)
        # print(visited)
        placeObjects(initialPos, (0,1), roomMap, visited)
    # print(len(visited))

if __name__ == '__main__':
    pt2()