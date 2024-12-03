import heapq
import os

# def pt1():
#     # print("rest")
#     # print(os.listdir() )
#     file = open("2024/day1/day1input.txt") 
#     line = file.readline()
#     left, right = [],[]
#     while line != '':
#         nums = line.split("   ")
#         # print(nums)
#         # print(nums[0].strip())
#         # print(nums[1].strip())
#         heapq.heappush(left, int(nums[0].strip()))
#         heapq.heappush(right, int(nums[1].strip()))
#         line = file.readline()
#     total = 0
#     while left and right:
#         total+= abs( heapq.heappop(left) - heapq.heappop(right))
#     print(total)
def main():
    # print("hi")
    file = open("2024/day1/day1input.txt")
    line = file.readline()
    left,right = [], {}
    while line != '':
        # print(line)
        nums = line.split("   ")
        l = int(nums[0].strip())
        left.append(l)
        r = int(nums[1].strip())
        if r in right:
            right[r]+= 1
        else:
            right[r] = 1
        line = file.readline()
    # print(left)
    # print(right)
    total = 0
    for num in left:
        # print("n=",num)
        if num in right:
            r = right[num]
            # print("r=",right[num])
            # print("l=", num)
            # print("adding", right[num] * num)
            total += (right[num] * num)
    print(total)
if __name__ == "__main__":
    main()
    # pt2()
