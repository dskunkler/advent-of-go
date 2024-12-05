from collections import defaultdict
def constructPageOrder(nums, pageOrder,parents):        
    parent, child = nums[0], nums[1]
    pageOrder[parent].append(child)
    parents[child].append(parent)
def generateValid(nums, valid, pageOrder ):
    s = set()
    print(nums)
    isValid = True
    for n in nums:
        if any(child in s for child in pageOrder[n]):
            isValid = False
            break
        s.add(n)
    if isValid:
        valid.append(nums)
def isValid(nums, pageOrder):
    s = set()
    print(nums)
    for n in nums:
        if any(child in s for child in pageOrder[n]):
            return False
        s.add(n)
    return True
def fixAndValidate(nums, valid, pageOrder):
    seen = set()
    pos = defaultdict(int)
    print(nums)
    isValid = True
    for i in range(len(nums)):
        if not isValid:
            break
        n = nums[i]
        for currNumsKids in pageOrder[n]:
            if currNumsKids in seen:
                parentLoc = pos[currNumsKids]
                newNums = nums[:parentLoc] + [n] + nums[parentLoc:i] + nums[i+1:]
                fixAndValidate(newNums, valid, pageOrder)
                isValid = False
                break
        pos[n] = i
        seen.add(n)
    if isValid:
        valid.append(nums) 
def addCenterNumber(valid):
    sol= 0
    for val in valid:
        sol += val[len(val)//2]
    return sol

def pt1():
    file = open("2024/day5/day5test.txt")
    line = file.readline()
    pageOrder, parents = defaultdict(list), defaultdict(list)

    while line != '\n':
        nums = [int(l) for l in line.split('|')]
        constructPageOrder(nums,pageOrder, parents)
        line = file.readline()
    line = file.readline()
    valid = []
    while line != "":
        nums = [int(l) for l in line.split(',')]
        generateValid(nums, valid, pageOrder)
        line = file.readline()

    print(addCenterNumber(valid))
                    
def pt2():
    file = open("2024/day5/day5.txt")
    line = file.readline()
    pageOrder, parents = defaultdict(list), defaultdict(list)

    while line != '\n':
        nums = [int(l) for l in line.split('|')]
        constructPageOrder(nums,pageOrder, parents)
        line = file.readline()
    line = file.readline()
    valid = []

    while line != "":
        nums = [int(l) for l in line.split(',')]
        if not isValid(nums, pageOrder): 
            print(nums, not valid)
            fixAndValidate(nums, valid, pageOrder)
        line = file.readline()
    print(addCenterNumber(valid))

if __name__ == '__main__':
    pt2()