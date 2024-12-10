def generateTrailHeads(trailMap):
    trailHeads = []
    for row in range(len(trailMap)):
        for col in range(len(trailMap)):
            if trailMap[row][col] == '0':
                trailHeads.append((row,col))
    return trailHeads

def findPeak(seeds, trailMap):
    seedRes = {}
    # canSeeFromPos = {}
    dirs = [(0,1), (1,0), (0,-1), (-1,0)]
    for row,col in seeds:
        seen = {(row,col)}
        # add to seen
        #search adjacent square
        # get how many can go to peak and save that as a result for that square
        
        seedRes[(row,col)] = 0
        print("searching", (row,col))
        for x,y in dirs:
            res = searchForHeigh((row+x, col+y), (row,col), trailMap, seen)
            # print("res from", (row +x,col + y), res)
            seedRes[(row,col)] += res
            # print(seedRes)
        print("res from ", (row,col), seedRes[(row,col)])
    return seedRes
        
        

def searchForHeigh(currPos, prevPos, trailMap, seen):
    row, col = currPos
    prevRow, prevCol = prevPos
    dirs = [(0,1), (1,0), (0,-1), (-1,0)]
    # print("curr", currPos)
    # print("prev", prevPos)
    if currPos in seen or not 0 <= row < len(trailMap) or not 0 <= col < len(trailMap[0]) or  trailMap[row][col] == '.' or int(trailMap[row][col]) - int(trailMap[prevRow][prevCol]) != 1:
        # print("oob or . or seen or not good value")
        print("leaving at ", (row,col))
        # seen.add(currPos)
        return 0
    # print("icon = ", trailMap[row][col])
    seen.add(currPos)
    if trailMap[row][col] == '9':
        print("found at", (row,col))
        # print("its 9")
        return 1
    canSee = 0

    for x,y in dirs:
        toAdd = searchForHeigh((row+x, col + y), currPos, trailMap, seen)
        canSee += toAdd
    # print("returning ", canSee)
    return canSee


    


def pt1():
    file = open("2024/day10/day10.txt")
    trailMap = [ list(line.strip()) for line in file.readlines()]
    # print(trailMap )
    trailHeads = generateTrailHeads(trailMap)
    print(trailHeads)
    res = findPeak(trailHeads, trailMap)
    print(res)
    sol = 0
    for key in res.keys():
        # print(key)
        # print(res[key])
        sol += res[key]
    print(sol)

if __name__ == '__main__':
    pt1()