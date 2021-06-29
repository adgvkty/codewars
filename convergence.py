from functools import reduce


def convergence(n):
    a = 1
    pos = 0
    nm = []
    am = []
    for i in range(10):
        am.append(a)
        nm.append(n)
        print(am)
        print(nm)
        if a == n:
            return pos
        else:
            pos += 1
        a = gen(a)
        n = gen(n)

        print(a, pos)

        print(n, pos)


def gen(nn):
    number = nn
    nn = list(str(nn))
    while '0' in nn:
        nn.remove('0')
    nn = [int(i) for i in nn]
    return reduce(lambda x, y: x * y, nn) + number


print(convergence(5))
