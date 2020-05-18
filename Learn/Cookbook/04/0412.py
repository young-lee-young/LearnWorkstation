# 在不同的容器进行迭代

from itertools import chain

name_list = ['liyao', 'zhangya', 'liuyafei', 'macheng']
age_list = (22, 23, 24, 25)


print(chain(name_list, age_list))
# <itertools.chain object at 0x0000020816F48908>


for info in chain(name_list, age_list): # 参数是可迭代对象
    print(info)


