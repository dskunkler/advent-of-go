def pt1():
    file = open("2024/day8/day8.txt")
    nodes = {'#' :[]}
    # go through each line and find the nodes, add their cordinates to the 
    nodeMap = [ line.strip() for line in file.readlines()]
    print(nodeMap)
    for row in range(len(nodeMap)):
        for col in range(len(nodeMap[0])):
            char = nodeMap[row][col]
            if char != '.':
                if char not in nodes:
                    nodes[char] = [(row,col)]
                else:
                    nodes[char].append((row,col))
    print(nodes)
    # go through each node pair, if the position of the anti node is in bounds, add it to nodes
    # [(0, 0), (0, 1), (0, 6), (0, 11), (1, 1), (1, 3), (2, 2), (2, 4), (2, 10), (3, 2), (3, 3), (4, 9), (5, 1), (5, 5), (5, 11), (6, 3), (6, 6), (7, 0), (7, 5), (7, 7), (8, 2), (9, 4), (10, 1), (10, 10), (11, 3), (11, 10), (11, 11)]
    for key in nodes.keys():
        if key == '#':
            continue
        pairs = nodes[key]
        print("looking at", key)
        for i in range(len(pairs)):
            for j in range(i+1,len(pairs)):
                print("first", pairs[i])
                print("second", pairs[j])
                nodes["#"].append(pairs[i])
                nodes["#"].append(pairs[j])
                firstRow, firstCol = pairs[i]
                secondRow, secondCol = pairs[j]
                ftsDif = (firstRow - secondRow, firstCol - secondCol)
                stfDif = (secondRow - firstRow, secondCol - firstCol)
                # print(ftsDif)
                it = 1
                ftsTarg = (firstRow + ftsDif[0], firstCol+ ftsDif[1])

                # print(stfDif)
                stfTarg = (secondRow + stfDif[0], secondCol+ stfDif[1])
                # print("stfTarg", stfTarg)
                while 0 <= stfTarg[0] < len(nodeMap) and 0 <= stfTarg[1] < len(nodeMap[0]):
                    print("in bounds", stfTarg)
                    nodes['#'].append(stfTarg)
                    it+=1 
                    stfTarg = (firstRow + it * stfDif[0], firstCol+ it * stfDif[1]) 
                    print("new stfTarg", stfTarg)
                it = 1
                while 0 <=ftsTarg[0] < len(nodeMap) and 0 <= ftsTarg[1] < len(nodeMap[0]):
                    print("in bounds", ftsTarg)
                    nodes['#'].append(ftsTarg)
                    it+=1
                    ftsTarg = (secondRow + it* ftsDif[0], secondCol+ it *ftsDif[1])
                    print("newFts targ", ftsTarg)
        print("sol", len(set(nodes['#'])))

if __name__ == '__main__':
    pt1()