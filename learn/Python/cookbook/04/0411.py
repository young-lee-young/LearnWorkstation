# 同时迭代多个序列

name_list = ['liyao', 'zhangya', 'liuyafei', 'macheng']
age_list = [22, 23, 24, 25]
for name, age in zip(name_list, age_list):  # 如果两个列表长度不同,以短的为基准
    print(name, age)
print('---------------------')


# 如果想以长的为基准,可以使用下面的方法
from itertools import zip_longest

del age_list[3]
for info in zip_longest(name_list, age_list):
    print(info)
print('---------------------')

# 长度不同,默认是用None来代替,可以以下方式添加默认值
for info in zip_longest(name_list, age_list, fillvalue=100):
    print(info)