# 用next方法迭代元素

import random
from collections import Iterable
from collections import Iterator

num_list = [random.randint(1,10) for i in range(10)]    # 列表是可迭代的对象
print(num_list)


print(isinstance(num_list, Iterable))   # 判断是否可以迭代
print(isinstance(num_list, Iterator))   # 列表不是迭代器


# while True:
#     next_num = next(num_list)   # 直接执行会报错,'list' object is not an iterator
#     print(next_num)


num_list_iter1 = iter(num_list) # 使用iter方法把列表,字典,字符串等变成迭代器
# 生成器可以迭代,生成器是迭代器对象
print(type(num_list_iter1))
print(isinstance(num_list_iter1, Iterator)) # 是迭代器


while True:
    try:
        next_num = next(num_list_iter1)
        print(next_num)
    except StopIteration:   # 使用next函数迭代,需要自己捕获异常
        print('迭代结束')
        break


num_list_iter2 = iter(num_list)
while True:
    next_num = next(num_list_iter2, None)
    print(next_num)
    if next_num == None:
        break
