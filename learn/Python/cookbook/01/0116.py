# 筛选序列中的元素

import random

num_list = [random.randint(-10, 10) for i in range(10)]


grt_0 = [n for n in num_list if n > 0]
print(grt_0)
let_0 = [n for n in num_list if n < 0]
print(let_0)


grt_0 = [n if n > 0 else 0 for n in num_list]   # 三元操作符
print(grt_0)
let_0 = [n if n < 0 else 0 for n in num_list]
print(let_0)


# 使用compress来判断是否符合条件
from itertools import compress
true_or_false = [n > 0 for n in num_list]   # 创建一个布尔序列
new_num_list = list(compress(num_list, true_or_false))    # 第一个参数是列表,第二个参数是一个布尔序列,返回的是可迭代的对象
print(new_num_list)


# filter函数使用,第一个参数是函数,第二个参数是列表,列表的每个数在函数中执行,符合条件返回True,不符合返回False
def is_greater0(num):
    if num > 0:
        return True
    else:
        return False

new_num_list = list(filter(is_greater0, num_list))  # filter返回的是可迭代的对象,用list强转
print(new_num_list)
print('------------------------')