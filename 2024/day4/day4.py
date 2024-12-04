import re

# def find_and_multiply(st):
#     return sum(int(a) * int(b) for a, b in re.findall(r'mul\((\d+),(\d+)\)', st))

def pt1():
    file = open("2024/day4/day4test.txt")
    line = file.readline()

def pt2():
    print("pt2")
    # with open("2024/day3/day3.txt") as file:
    #     line = file.readline()

    # instructions = re.split(r"(don't\(\))|(do\(\))", line)
    # total = 0
    # multiply = True

    # for segment in instructions:
    #     if segment == "don't()":
    #         multiply = False
    #     elif segment == "do()":
    #         multiply = True
    #     elif segment and multiply:
    #         total += find_and_multiply(segment)

    # print(total)


if __name__ == '__main__':
    pt1()