with open("day3.txt") as f:
    lines = f.readlines()

def check(right, down, currentLine=0, currentHoriz=0, count=0):
    try:
        nextLine = lines[currentLine+down]
        currentLine += down
    except IndexError:
        #Reached the bottom
        return count
    
    
    #Check along $right down $down
    #Don't run off the end - it repeats
    currentHoriz += right

    if nextLine[currentHoriz % (len(nextLine)-1)] == "#":
        count += 1
    return check(right, down, currentLine, currentHoriz, count)


#Go
slopes = [[1,1], [3,1], [5,1], [7,1], [1,2]]

out = 1
for slope in slopes:
    out = out * check(slope[0],slope[1])

print(out)