def pt1():
    file = open("2024/day4/day4.txt")
    line = file.readline()
    wordMap = []
    while line != '':
        wordMap.append(line.strip())
        line = file.readline()
    sol= 0
    for row in range(len(wordMap)):
        for col in range(len(wordMap[0])):
            word =wordMap[row][col:col+4] 
            if word == 'XMAS' or word == 'SAMX':
                sol+=1
            if row <= len(wordMap) - 4:
                word =wordMap[row][col] +  wordMap[row+1][col] + wordMap[row+2][col] + wordMap[row+3][col]
                if word == 'XMAS' or word == 'SAMX':
                    sol+=1
                if col  <= len(wordMap[0]) - 4:
                    word =wordMap[row][col] +  wordMap[row+1][col+1] + wordMap[row+2][col+2] + wordMap[row+3][col+3]
                    dl =wordMap[row+3][col] +  wordMap[row+2][col+1] + wordMap[row+1][col+2] + wordMap[row][col+3] 
                    if word == 'XMAS' or word == 'SAMX':
                        sol +=1
                    if dl == 'XMAS' or dl == 'SAMX':
                        sol +=1
    print(sol)
                    
def pt2():
    file = open("2024/day4/day4.txt")
    line = file.readline()
    wordMap = []
    while line != '':
        wordMap.append(line.strip())
        line = file.readline()
    sol= 0
    for row in range(len(wordMap)):
        for col in range(len(wordMap[0])):
            if row <= len(wordMap) - 3 and col  <= len(wordMap[0]) - 3:
                words = set(['SAM', 'MAS'])
                word =wordMap[row][col] +  wordMap[row+1][col+1] + wordMap[row+2][col+2] 
                dl =wordMap[row+2][col] +  wordMap[row+1][col+1] + wordMap[row][col+2]
                if word in words and dl in words:
                    sol +=1
    print(sol)


if __name__ == '__main__':
    pt2()