import sys

pythonlist = [[0,0,0],[0,0,0],[0,0,0]] #The "slice" - a 2D list in python.

def addTwo(): #Makes sure the top left item is a 2.
    try: #lazy error handling in case you shrink the list to [] i.e nothing.
        pythonlist[0][0] = 2 
    except:
        return

def printmap(): #Prints out the map
    for index, i in enumerate(pythonlist):
        for anotherindex, b in enumerate(pythonlist[index]):
            if pythonlist[index][anotherindex] == 2: #Makes sure there's a 2 in the list.
                for i in pythonlist:
                    print(i)
                return
    addTwo() #if not, we had better add a 2 to the list.
    for i in pythonlist:
        print(i)
    return

def move(x2): #takes the input and depending on what it is, does some stuff.
    global x
    if len(pythonlist) == 1:
            return
    if x2 == "w":
        for yindex, i in enumerate(pythonlist):
            for xindex, a in enumerate(pythonlist[yindex]): #moves the 2 up.
                if pythonlist[yindex][xindex] == 2:
                    pythonlist[yindex][xindex] = 0
                    pythonlist[yindex-1][xindex] = 2
                    printmap()
                    x = 0
                    return
    if x2 == "s":
        for yindex, i in enumerate(pythonlist):
            for xindex, a in enumerate(pythonlist[yindex]):#moves the 2 down.
                if pythonlist[yindex-1][xindex] == 2:
                    pythonlist[yindex-1][xindex] = 0
                    pythonlist[yindex][xindex] = 2
                    printmap()
                    x = 0
                    return
    if x2 == "a":
        for yindex, i in enumerate(pythonlist):
            for xindex, a in enumerate(pythonlist[yindex]): #moves the 2 left.
                if pythonlist[yindex][xindex] == 2:
                    pythonlist[yindex][xindex-1] = 2
                    pythonlist[yindex][xindex] = 0
                    printmap()
                    x = 0
                    return
    if x2 == "d":
        for yindex, i in enumerate(pythonlist):
            for xindex, a in enumerate(pythonlist[yindex]): #moves the 2 right.
                if pythonlist[yindex][xindex-1] == 2:
                    pythonlist[yindex][xindex] = 2
                    pythonlist[yindex][xindex-1] = 0
                    printmap()
                    x = 0
                    return

def size(x):
    if x == "gr":
        for index, i in enumerate(pythonlist): #grows the list by list[len(list[n])+1][len(list[n][n])+1]
            pythonlist[index].append(0)
        pythonlist.append([])
        for i in range(len(pythonlist)):
            pythonlist[-1].append(0)
        printmap()
        x = 0
        return
    if x == "sh":
        for index, i in enumerate(pythonlist): #shrinks the list by list[len(list[n])-1][len(list[n][n])-1]
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

