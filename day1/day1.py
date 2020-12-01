
with open("day1.txt") as f:
    lines = f.readlines()

#Convert all to integers
lines = list(map(int, lines))

#Part 1 of puzzle
found = False
for line in lines:
    for line2 in lines:
        if line + line2 == 2020:
            found = True
            correctOne = line
            correctTwo = line2
    if found:
        break

print(correctOne*correctTwo)

#Part 2 of puzzle
found = False
for line in lines:
    for line2 in lines:
        for line3 in lines:
            if line + line2 + line3 == 2020:
                found = True
                correctOne = line
                correctTwo = line2
                correctThree = line3
    if found:
        break

print(correctOne*correctTwo*correctThree)