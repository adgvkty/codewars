def is_a_valid_message(message):
    pos = 0
    if message == '':
        return True
    elif '0' <= message[pos] <= '9':
        for i in range(int(message[pos])):
            if message[i+1].isnumeric():
                return False
    else:
        return False

def check_letters(message, pos):
    for i in range(int(message[pos])):
        if message[i + 1].isnumeric():
            return False


print(is_a_valid_message('4cod13hellocodewars'))