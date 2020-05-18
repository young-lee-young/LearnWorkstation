# 分组

from itertools import groupby
from operator import itemgetter

info_dict = [
    {'name': 'liyao', 'age': 25, 'sex': 'boy'},
    {'name': 'zhangya', 'age': 23, 'sex': 'girl'},
    {'name': 'liuyafei', 'age': 22, 'sex': 'boy'},
    {'name': 'macheng', 'age': 21, 'sex': 'boy'}
]

info_dict.sort(key=itemgetter('sex'))
for sex, items in groupby(info_dict, key=itemgetter('sex')):    # groupby返回的迭代器
    # 迭代时返回的是一个值和一个子迭代器
    print(sex, items)
    for i in items:
        print('', i)
print('------------------------')