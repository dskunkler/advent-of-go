import re

# We want to take the cost and the x,y,z and 
# Should we just dfs it?
def findPrize(a,b, prize):
    ax,ay, aCost =a
    bx, by, bCost = b
    prizeX, prizeY = prize
    maxAPresses = max(prizeX //ax, prizeY // ay)
    maxBPresses = max(prizeX //bx, prizeY // by)

    maxCost = 2 * max(maxAPresses, maxBPresses) * aCost
    minCost = maxCost

    for i in range(maxAPresses+1):
        for j in range(maxBPresses+1):
            currX = ax * i + bx * j
            currY = ay * i + by * j
            cost = i* aCost + j*bCost
            if currX == prizeX and prizeY == currY:
                minCost = min(minCost, cost)
            
    return 0 if minCost == maxCost else minCost

def findSmartPrize(a,b,prize):
    ax,ay, aCost = a
    bx, by, bCost = b
    prizeX, prizeY = prize

    numerator = prizeX* ay - ax*prizeY
    denominator = bx*ay - ax*by
    if denominator == 0 or numerator%denominator != 0:
        return 0
    j = numerator // denominator
    numerator = prizeX - j* bx
    denominator = ax
    if ax <=0 or numerator %ax != 0:
        return 0
    i = numerator// denominator
    return aCost * i + bCost * j
    

#22262
def pt1():
    file = open("2024/day13/day13.txt")
    lines = [line.strip() for line in file.readlines()]
    sol = 0
    for i in range(0,len(lines),4):
        a = [ int(i) for i in re.findall("(\d+)", lines[i])]
        b = [int(i) for i in re.findall("(\d+)", lines[i+1])]
        prize = [int(i) for i in re.findall("(\d+)", lines[i+2])]

        res = findSmartPrize( (a[0], a[1], 3), (b[0], b[1], 1), ( 10000000000000+prize[0],10000000000000+ prize[1]))
        sol +=res
    print(sol)

def pt2():
    file = open("2024/day13/day13.txt")

if __name__ == '__main__':
    pt1()