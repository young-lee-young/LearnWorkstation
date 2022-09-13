# -*- encoding: utf-8 -*-

# range、random函数的使用

import random


# range函数的使用
# 还有一个xrange函数,在python3里被弃用,返回的是一个生成器对象,不能切片
range_num = range(10)
print(range_num)
# range(0, 10)
print(type(range_num))
# <class 'range'>


# 隔一个取一个
for num in range(0, 10, 2):
    print(num)

# 从大到小取
for num in range(10, 0, -1):
    print(num)


# 产生0-1之间的浮点数
random_float = random.random()
print('随机0-1浮点数:\n', random_float)


# 产生两个数之间的整数,左右都是闭
random_int2 = random.randint(4, 10)
print('随机整数:\n', random_int2)


# 一次生成多个整数随机数
data = [random.randint(1, 10) for _ in range(10)]
print('一次生成多个随机数:\n', data)


# 正态分布
normalvariate = [random.normalvariate(0,1) for _ in range(100)]
print('随机正态分布:\n',normalvariate)


# 在一个序列里随机选择一个
name_list = ['liyao', 'macheng', 'liwei', 'liangbo', 'liuyafei', 'xuwei', 'pengzeyu']
name = random.choice(name_list)
print(u'随机选择:\n', name)  # xuwei


all_string = 'liyaozhangya'
one_char = random.choice(all_string)
print('随机选择:\n', one_char)  # y