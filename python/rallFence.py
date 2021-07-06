def encode_rail_fence_cipher(string, n):
    array = list(string)
    result = ''
    
    for i in range(n):
        result = result + string[::n+n//2]
    
    return result

def decode_rail_fence_cipher(string, n):    
    if string == '':
        return '' 
    

print(encode_rail_fence_cipher('WEAREDISCOVEREDFLEEATONCE', 3))

'''
WEAREDISCOVEREDFLEEATONCE

WECRLTEERDSOEEFEAOCAIVDEN
WECRLTEERDSOEEFEAOCAREDISCOVEREDFLEEATONCE

H !e,Wdloollr

'''
