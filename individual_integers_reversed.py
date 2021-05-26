import re

def solution(args):
    result = []
    strings = re.split(r',', args)
    for el in strings:
        if len(el) > 2:
            el = re.split(r'-', el)
            if el[0] == '':
                _, start, end = el
                start = -int(start)
            else:
                start, end = el
            for i in range(int(start), int(end)):
                result.append(i)
        else:
            result.append(int(el))

    return result


string = '-6,-3-1,3-5,7-11,14,15,17-20'
print(solution(string))