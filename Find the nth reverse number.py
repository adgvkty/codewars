def find_reverse_number(n):
    i = 0
    count = 0
    while count < n:
        if int(str(i)[::-1]) == i:
            count += 1
        i += 1
    return i - 1


print(find_reverse_number(1000000))