import random

a = 1
b = 850

def solve(a, b):
    temp = []
    dct = {}
    res = []
    for i in range(a, b):
        for x in range(1, i+1):
            if i % x == 0:
                temp.append(x)
        key = sum(temp) / i
        temp = []
        dct.setdefault(key, [])
        dct[key].append(i)
    dct = list(dct.items())
    for i in range(len(dct)):
        if len(dct[i][1]) >= 2:
            res.append(min(dct[i][1]))
    print(sum(res))

solve(a, b)