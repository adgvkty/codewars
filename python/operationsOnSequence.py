from functools import reduce
import math
import time


def solve(arr):
    a = []
    i = 0
    while i < len(arr):
        a.append((arr[i] * arr[i]) + (arr[i + 1] * arr[i + 1]))
        i = i + 2
    a = reduce(lambda x, y: x * y, a)
    print(a)
    if a == 0:
        return [0, 0]
    return find_square(a)


def find_square(n):
    t0 = time.time()
    s = dict()
    for i in range(n):
        if i * i > n:
            break
        s[i * i] = 1
        if (n - i * i) in s.keys():
            print(time.time() - t0)
            return [int((n - i * i) ** (1 / 2)), i]
    return [0, 0]


print(solve([3, 9, 8, 4, 6, 8, 7, 8, 4, 8, 5, 6, 6, 4, 4, 5]))
