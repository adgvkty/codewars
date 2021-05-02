import math


def solve(a, b):
    count = 0
    old_i = 0
    x = 0
    for i in range(a, b):
        if i == 1:
            continue
        else:
            old_i = gen(i)
            i = gen(i)
            while x < 1:
                i = gen(i)
                if i == 1:
                    count += 1
                    x = 1
                elif old_i == i:
                    x = 1
    return count


def gen(p):
    if p < 10:
        return p**2
    else:
        p = list(str(p))
        while '0' in p:
            p.remove('0')
        p = list(map(int, p))
        for i in range(len(p)):
            p[i] = p[i]**2
        return sum(p)


print(solve(1, 25))