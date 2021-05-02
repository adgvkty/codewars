def match(s):
    s = list(s)
    d = {'A':'B', 'B':'A', 'X':'Y', 'Y':'X'}
    c = []
    print(s)
    res = 0
    for i in range(len(s)):
        try:
            c.append([s.index(s[i]), s.index(d.get(s[i]))])
            print(c)
            print(s)
            res += 1
        except Exception as e:
            print(e)
    return res // 2

print(match('BXABAYBA'))