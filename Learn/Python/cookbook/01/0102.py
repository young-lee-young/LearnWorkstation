# 从任意长度的可迭代对象中分解元素

import random

num_list = [random.randint(1,100) for i in range(100)]
print(num_list)


num1, *num2, num3 = num_list  # num2是除第一和最后外的值
print(num2)
print('------------------------')