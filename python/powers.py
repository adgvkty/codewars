def last_digit(lst):
    if len(lst) == 2:
        a, b = lst
        return a**b
    elif len(lst) == 0 or len(lst) == 1:
        return 1
    else:
        a, b, c = lst
        b_c = power(b, c)
        return a**b_c
    
def power(base, exp):
    print('calculation power')
    if exp < 0:
        return 1 / power(base, -exp)
    ans = 1
    while exp:
        if exp & 1:
            ans *= base
        exp >>= 1
        base *= base
    return ans

print(last_digit([32, 4, 22]))