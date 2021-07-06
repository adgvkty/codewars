badSymbols = ['/', ',', '.', '!', '?', ':', ';', '-']

res = {}

def top_3_words(text):
    for symbol in badSymbols:
        text = text.replace(symbol, '')
    text = [word.lower() for word in text.replace('\n', ' ').split(' ')]
    
    for word in text:
        if word in res.keys():
            res[word] += 1
        else:
            res.setdefault(word, 1)
    
    firstBest = ['', 0]
    secondBest = ['', 0]
    thirdBest = ['', 0]
    
    for item in res.items():
        if item[0] == '':
            continue
        if item[1] > firstBest[1]:
            thirdBest[0], thirdBest[1] = secondBest[0], secondBest[1]
            secondBest[0], secondBest[1] = firstBest[0], firstBest[1]
            firstBest[0], firstBest[1] = item[0], item[1]
        elif item[1] > secondBest[1]:
            thirdBest[0], thirdBest[1] = secondBest[0], secondBest[1]
            secondBest[0], secondBest[1] = item[0], item[1]
        elif item[1] > thirdBest[1]:
            thirdBest[0], thirdBest[1] = item[0], item[1]
    
    
    resStack = [firstBest[0], secondBest[0], thirdBest[0]]
    for item in resStack:
        try:
            resStack.remove('')
            if len(resStack) == 1 and resStack[0] == '':
                resStack = []
        except:
            continue

    return resStack

print(top_3_words("e e e e DDD ddd DdD: ddd ddd aa aA Aa, bb cc cC e e e"))