import time


def dup(arry):
    t = ''
    for i in range(len(arry)):
        for s in arry[i]:
            if s in t:
                pass
            else:
                t = t + s
            print("With Order:", t)
            arry[i] = t
            t = ''

    return arry


print(dup(["ccooddddddewwwaaaaarrrrsssss", "piccaninny", "hubbubbubboo"]))
