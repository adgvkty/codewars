import time


def dbl_linear(n):
    u = [1]
    for i in range(n*2):
        u.append((2*u[i])+1)
        u.append((3*u[i])+1)
    print(u)
    u = set(u)
    u = list(u)
    u.sort()
    return u[n]


print(dbl_linear(500))
