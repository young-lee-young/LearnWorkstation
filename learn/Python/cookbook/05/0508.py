# 以块的模式进行迭代

from functools import partial

size = 100  # 块的大小

with open('D:/spider/okgaokaozhiyuan.py', 'rt', encoding='utf-8') as f:
    records = iter(partial(f.read, size), r'')
    for i in records:
        print(i)
        print('-------------------')