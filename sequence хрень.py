from functools import reduce


def convergence(n):
    b = 1
    x = 0
    base = []
    pos = 0
    for i in range(5000):
        base.append(b)
        b = gen(b)
    while x < 1:
        if n in base:
            return pos
        else:
            n = gen(n)
            pos += 1


def gen(n):
    base_number = n
    n = list(str(n))
    while '0' in n:
        n.remove('0')
    n = [int(i) for i in n]
    n = int(reduce(lambda x, y: x * y, n)) + int(base_number)
    return n


print(convergence(500))
