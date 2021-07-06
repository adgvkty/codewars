import re


def get_consective_items(items, key):
    items = str(items)
    key = str(key)
    return len(max(re.findall(rf'[{key}]+', items)))


print(get_consective_items('ascasdaiiiasdacasdiiiiicasdasdiiiiiiiiiiisdasdasdiii', 'i'))
