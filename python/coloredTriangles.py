def triangle(row):
    i = 0
    d = {'GG': 'G', 'BG': 'R', 'RG': 'B', 'BR': 'G'}
    new_row = []
    while i + 1 < len(row):
        new_row.append(d.get(row[i]+row[i+1]))
        print(d.get(row[i]+row[i+1]), row[i]+row[i+1])
        i += 1

    return new_row


print(triangle('RBRGBRBGGRRRBGBBBGG'))
