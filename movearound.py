import sys

pythonlist = [[0,0,0],[0,0,0],[0,0,0]]

def addTwo():
    try:
        pythonlist[0][0] = 2
    except:
        return

def printmap():
    for index, i in enumerate(pythonlist):
        for anotherindex, b in enumerate(pythonlist[index]):
            if pythonlist[index][anotherindex] == 2:
                for i in pythonlist:
                    print(i)
                return
    addTwo()
    for i in pythonlist:
        print(i)
    return

def move(x2):
    global x
    if len(pythonlist) == 1:
            return
    if x2 == "w":
        for yindex, i in enumerate(pythonlist):
            for xindex, a in enumerate(pythonlist[yindex]):
                if pythonlist[yindex][xindex] == 2:
                    pythonlist[yindex][xindex] = 0
                    pythonlist[yindex-1][xindex] = 2
                    printmap()
                    x = 0
                    return
    if x2 == "s":
        for yindex, i in enumerate(pythonlist):
            for xindex, a in enumerate(pythonlist[yindex]):
                if pythonlist[yindex-1][xindex] == 2:
                    pythonlist[yindex-1][xindex] = 0
                    pythonlist[yindex][xindex] = 2
                    printmap()
                    x = 0
                    return
    if x2 == "a":
        for yindex, i in enumerate(pythonlist):
            for xindex, a in enumerate(pythonlist[yindex]):
                if pythonlist[yindex][xindex] == 2:
                    pythonlist[yindex][xindex-1] = 2
                    pythonlist[yindex][xindex] = 0
                    printmap()
                    x = 0
                    return
    if x2 == "d":
        for yindex, i in enumerate(pythonlist):
            for xindex, a in enumerate(pythonlist[yindex]):
                if pythonlist[yindex][xindex-1] == 2:
                    pythonlist[yindex][xindex] = 2
                    pythonlist[yindex][xindex-1] = 0
                    printmap()
                    x = 0
                    return

def size(x):
    if x == "gr":
        for index, i in enumerate(pythonlist):
            pythonlist[index].append(0)
        pythonlist.append([])
        for i in range(len(pythonlist)):
            pythonlist[-1].append(0)
        printmap()
        x = 0
        return
    if x == "sh":
        for index, i in enumerate(pythonlist):
            pythonlist[index].pop()
        pythonlist.pop()
        printmap()
        x = 0
        return

def start():
    global x
    print("Press w, s, a, d to move up/down/left/right")
    print("Press q to quit.")
    addTwo()
    quitgame = 0
    while quitgame == 0:
        x = input("Type here: ")
        if x == "w":
            move(x)
        if x == "s":
            move(x)
        if x == "a":
            move(x)
        if x == "d":
            move(x)
        if x == "sh":
            size(x)
        if x == "gr":
            size(x)
        if x == "q":
            quitgame = 1

start()

