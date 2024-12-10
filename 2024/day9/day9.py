def generateIdLine(nums):
    it, ind = 0,0
    res = []
    while it < len(nums):
        block = int(nums[it])
        for i in range( block):
            res.append(ind)
        if it+1 < len(nums):
            skip = int(nums[it+1])
            for i in range(skip):
                res.append('.')
        # print("new res", res)
        it+=2
        ind+=1
    return res

def swapStuff(idLine):
    l, r = 0, len(idLine)-1
    idLine = idLine
    # print(idLine)
    while l < r:
        # scan for next valid
        while l < r and idLine[l] != '.':
            l+=1
        while l < r and idLine[r] == '.':
            r-=1
        if l < r:
            idLine[l], idLine[r] = idLine[r], '.'
            # print("new idline", idLine)
    return idLine
def swapComplicated(idLine):
    l, r = 0, len(idLine)-1
    while l < r:
        while l < r and idLine[r] == '.':
            r-=1
        rl = r
        while rl > l and idLine[r] == idLine[rl]:
            rl-=1
            # print("rl now seeing", idLine[rl])
        # print(idLine[rl+1], "len", r - rl)
        spaceNeeded = r-rl
        # point it bacK To the right spot...?
        rl+=1
        # scan for next valid
        while l < r and idLine[l] != '.':
            l+=1
        lr = l
        fit = False
        while lr < rl and not fit:
            currL = lr            
            # print("lr =", lr)
            # print("rl = ", rl)
            #find the amount of space
            while lr  < rl and idLine[lr] == '.':
                lr+=1
            freeSpace = lr- currL
            # print("free space", freeSpace)
            if freeSpace >= spaceNeeded:
                # print("it fits")
                # print("moving from ", rl, "to", currL)
                for i in range(spaceNeeded):
                    idLine[currL+i], idLine[rl+i] = idLine[rl+i], '.'
                fit = True
                # print("lr now", lr)
            else:
                # find next free space and start again
                lr+=1
                while lr< rl and idLine[lr]!= '.':
                    lr+=1
        # print(idLine)
        if not fit:
            # print("couldn't fit it")
            r = rl -1
        # if l < r:
        #     idLine[l], idLine[r] = idLine[r], '.'
        #     # print("new idline", idLine)
    return idLine

def checksum(idLine):
    sol = 0
    for i in range(len(idLine)):
        if idLine[i] == '.':
            continue
        sol += int(idLine[i]) * i
        # print("int(idline)", int(idLine[i]))
        # print("i=", i)
    return sol

#90786863565      

def pt1():
    file = open("2024/day9/day9.txt")
    input = ""
    for line in file.readlines():
        input += line.strip()
    # print(input)
    idLine = generateIdLine(input)
    # print("generated", idLine)
    idLine = swapComplicated(idLine)
    # print("swapped", idLine)
    # print("swapped", idLine)
    print(checksum(idLine))

if __name__ == '__main__':
    pt1()