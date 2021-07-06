import re

s = '~c0d3w4rs~'
symb = ['?', '.', '!', '/', "'", ';', ',', '"','-',' ',':','~']


def licence_plate(s):
    s = list(s)
    for i in range(len(s)):
        if s[i] in symb:
            s[i] = ''
    s = [v.upper() for v in s if v]
    s = ''.join(s)
    for i in range(len(s)-1):
        if s[i].isdigit() and s[i+1].isalpha() or s[i].isalpha() and s[i+1].isdigit():
            s = s[:i+1] + '-' + s[i+1:]
    s = s[:8]
    if s[len(s)-1] == '-':
        return s[:7:]
    elif s[0] == '-':
        return s[1:]
    else:
        return s


def old(s):
    print(s)
    s = re.split(r'([\d, ,\',\,,_,-,\*]+)', s)
    print(s)
    s = (''.join(e.upper() for e in el if e.isalnum()) for el in s)
    s = '-'.join(s)
    s = s[:8:]
    s = s.replace(' ', '-')
    if s.isdigit() or len(s) == 1:
        return 'not possible'
    if s[len(s)-1] == '-':
        return s[:7:]
    elif s[0] == '-':
        return s[1:]
    else:
        return s


print(licence_plate(s))
    