from collections import defaultdict
def applyRules(stones):
    newStones = []
    for stone in stones:
        if stone == '0':
            newStones.append('1')
        elif len(stone) % 2 == 0:
            mid = len(stone)//2
            newStones.append(str(int(stone[:mid])))
            newStones.append(str(int(stone[mid:])))
        else:
            newStones.append(str(int(stone) * 2024))
    return newStones
def blink(times, stones):
    # print("initial")
    for _ in range(times):
        stones = applyRules(stones)
    return len(stones)

def pt1():
    file = open("day11.txt")
    stones = file.readline().strip().split()
    print(stones)
    print(blink(75, stones))

def applyDyanmicRule(stone):
    newStones = []
    if stone == 0:
        newStones.append('1')
    elif len(stone) % 2 == 0:
        mid = len(stone)//2
        newStones.append(str(int(stone[:mid])))
        newStones.append(str(int(stone[mid:])))
    else:
        newStones.append(str(int(stone) * 2024))
    return newStones

def dynamic(times, stones, cache):
    if times == 0:
        return stones
    
    res = []
    curr = []
    for stone in stones:
        if (times, stone) in cache:
            res = cache[(times, stone)]
            # print("ress from cache", res)
        else:
            newStones = applyDyanmicRule(stone)
            # print("after transform", newStones)
            whole = []
            for newStone in newStones:
                curr = dynamic(times-1, [newStone], cache)
                res = res + curr
            # print("res to cache for ",(times, stone), res)
            cache[(times, stone)] = res
        # print("res", res)
    return res

def blinkDyanmic(times, stones):
    # print("seraching", (times, stones))
    cache = {}
    res = []
    for stone in stones:
        # print("searching stone", stone)
        # apply the rule times amount of times. ONCE we get to a solution, cache what every step creates.
        currRes = dynamic(times, [stone], cache)
        # print("blink res..", currRes)
        res = res+currRes
    # print("final cache", cache)
    return res


def pt2():
    file = open("day11.txt")
    stones = file.readline().strip().split()
    print(stones)
    # sol = blinkDyanmic(50, stones)
    counter = {}
    for stone in stones:
        counter[stone] = counter.get(stone, 0) + 1

    print(counter)

    for i in range(75):
        newCt = {}
        for key,val in counter.items():
            if key == '0':
                newCt['1'] = newCt.get('1',0) + val
            elif len(key)%2 == 0:
                n = len(key)//2
                lhs = str(int(key[:n]))
                rhs = str(int(key[n:]))
                newCt[lhs] = newCt.get(lhs, 0) + val
                newCt[rhs] = newCt.get(rhs, 0) + val
            else:
                newVal = str(int(key) * 2024)
                newCt[newVal] = newCt.get(newVal,0) + val
        counter = newCt
        print(counter)

    sol = 0
    for key,val in counter.items():
        sol += val
    print("final sol", sol)
    # print("sol", len(sol))


if __name__ == '__main__':
    pt2()