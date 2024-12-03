def pt1():
    file = open("2024/day2/day2.txt") 
    line = file.readline()
    sol = 0

    while line != '':
        # print(line)
        nums = [int(x) for x in line.split(' ')]
        print(nums)
        line = file.readline()
        valid = True
        l, r = nums[0], nums[1]
        #decreasing
        available = 1
        if r < l:

            for i in range(1, len(nums)):
                l,r = nums[i-1], nums[i]
                dif = l-r
                print("dif=", dif)
                if dif > 3:
                    print("invalid")
                    valid = False
                    break
                if l == r:
                    if available > 1  and i == len(nums) or i < len(nums)-1 and nums[i+1] != r:
                        available = 0
                        continue
                    else:
                        valid = False
                        break
            sol += 1 if valid else 0
            print("new sol",sol )
        #increasing
        elif l < r:
            for i in range(1, len(nums)):
                l,r = nums[i-1], nums[i]
                dif = r-l
                print("dif=", dif)
                if dif > 3 :
                    print("invalid")
                    valid = False
                    break
                if l == r:
                    print("equal")
                    if available > 1  and i == len(nums) or i < len(nums)-1 and nums[i+1] != r:
                        available = 0
                        continue
                    else:
                        valid = False
                        break
            sol += 1 if valid else 0
            print("new sol", sol)
    print(sol)

def pt2():
    file = open("2024/day2/day2.txt") 
    line = file.readline()
    sol = 0
    def helper(nums, swaps):
    # Determine if the current list is strictly increasing or decreasing
        def is_valid(nums):
            increasing = all( 1 <= nums[i] - nums[i-1] <= 3 for i in range(1, len(nums)))
            decreasing = all(1<= nums[i-1] - nums[i] <= 3 for i in range(1, len(nums)))
            return increasing or decreasing

        # Base Case: if valid sequence
        if is_valid(nums):
            return True
        
        # Invalid case and swaps still exist
        for i in range(len(nums)):
            if swaps > 0:
                # Remove element i attempt fixed safe run.
                if helper(nums[:i] + nums[i+1:], swaps - 1):
                    return True
        return False  # Out of Swaps 
                        
    while line != '':
        nums = [int(x) for x in line.split(' ')]
        print(nums)
        line = file.readline()
        if helper(nums, 1):
            print("good")
            sol+=1
        print(sol)

if __name__ == '__main__':
    pt2()
