import re

def is_a_valid_message(message):
    digits = re.findall(r'\d+', message)
    chars = re.findall(r'\D+', message)
    try:
        for i in range(len(chars)):
            if int(digits[i]) == len(chars[i]):
                continue
            else:
                return False
    except Exception:
        return False
    return True

print(is_a_valid_message('3hey5hello2hi5'))