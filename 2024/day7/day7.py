from collections import deque
def canEqual(targ, options):
    print("targ = ",targ)
    print("optinos = ", options)

    curr = (options.popleft(), options)
    q = deque()
    q.append(curr)
    while q:
        if not q:
            break
        val, opts = q.popleft()
        # print("val, opts", val, opts)
        # print(targ == val)
        if not opts:
            # print("no opts")
            if val == targ:
                print("true returning")
                return True
            else:
                continue
        next = opts.popleft()
        # print("next, opts", next, deque(opts))
        q.append((val+next, deque(opts)))
        # print("adding", val+next)
        q.append((val*next, deque(opts)))
        q.append((int(str(val)+str(next)), deque(opts)))

    return False    
        
def pt1():
    file = open("2024/day7/day7.txt")
    vals = []
    for line in file.readlines():
        nums = line.split(':')
        vals.append((int(nums[0]), deque([int(n) for n in nums[1].strip().split()])))
    sol = 0
    for targ, options in vals:
        if canEqual(targ, options):
            print("can equal")
            sol += targ
            print("new sol", sol)
    print("sol", sol)
                    
def pt2():
    file = open("2024/day7/day7.txt")

if __name__ == '__main__':
    pt1()