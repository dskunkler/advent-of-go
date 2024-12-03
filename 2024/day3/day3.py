import re
def pt1():
    file = open("2024/day3/day3test.txt")
    line = file.readline()

    sol = findAndMul(line)
    print(sol)

def findAndMul(st):
    muls = re.findall('mul\((\\d+),(\\d+)\)', st)
    # print(muls)
    sol = 0
    for a,b in muls:
        sol+= int(a)*int(b)
    # print(sol)
    return sol


def pt2():    
    file = open("2024/day3/day3.txt")
    line = file.readline()

    # print(line)
    inst = re.split("(don't\(\))|(do\(\))", line)
    # print(inst)
    sol = 0
    mul = True
    for st in inst:
        # print(st)
        if st == "don't()":
            mul = False
            
        elif st == "do()":
            mul = True
        elif st and mul:
            sol += findAndMul(st)
    print(sol)



if __name__ == '__main__':
    pt2()