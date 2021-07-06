import re

def solution(args):
    result = ''
    tmp = []
    for i in range(len(args)):
        
        if i + 1 < len(args) and args[i] + 1 == args[i+1]:
            tmp.append(args[i])
        elif len(tmp) >= 3:
            start, *c, end = tmp
            result = result + ',' + str(start) + '-' + str(args[i])
            tmp = []
        elif len(tmp) == 2 and args[i-1] == args[i] - 1:
            start, *c, end = tmp
            result = result + ',' + str(start) + '-' + str(args[i])
            tmp = []
        else:
            if tmp:
                result = result + ',' + str(tmp.pop()) + ',' + str(args[i])
                tmp = []
            else:
                result = result + ',' + str(args[i])
    return result[1:]

array = [-6,-3,-2,-1,0,1,3,4,5,7,8,9,10,11,14,15,17,18,19,20]
print(solution(array))

'''
Answer:
-6,-3-1,3-5,7-11,14,15,17-20
Result:
-6,-3-1,3-5,7-11,14,15,17-20
'''