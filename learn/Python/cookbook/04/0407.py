# 对迭代器进行切片操作

from collections import Iterator
from itertools import islice

def count(n):
    while True:
        yield n
        n += 1


c = count(1)
print(isinstance(c, Iterator))


# cut = c[10:20] # 这种切片方法不正确


cut = islice(c, 10, 20)
print(cut)
for i in cut:
    print(i)