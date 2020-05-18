# 扁平化处理嵌套类型的序列

from collections import Iterable

def flatten(items, ignore_types=(str, bytes)):
    for x in items:
        if isinstance(x, Iterable) and not isinstance(x, ignore_types):
            yield from flatten(x)   # 注意yield from的使用
        else:
            yield x


items = [1, 2, [3, 4, [5, 6, 7], 8, 9], 10]


for x in flatten(items):
    print(x)