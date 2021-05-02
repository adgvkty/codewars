def pairs(a):
    count = 0
    i = 0
    while i + 1 < len(a):
        if abs(abs(a[i] - a[i + 1])) == 1:
            count += 1
        i += 2
    return count


print(pairs([21, 20, 22, 40, 39, -56, 30, -55, 95, 94]))
