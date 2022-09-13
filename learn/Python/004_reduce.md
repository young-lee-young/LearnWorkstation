# reduce函数用法

reduce函数已经被移动到functools下
第一个参数是函数名，第二个参数是可迭代对象，可以有第三个初始值参数
以迭代对象前两项执行函数，得到的结果与第三项执行函数


```python
from functools import reduce

num_list = range(1, 5)
# num_list: [1, 2, 3, 4]

def func_plus(x, y):
    return x + y
    
num_result = reduce(func_plus, num_list)
# type(num_result): int
# new_result: 10

# 有初始值参数
new_result_2 = reduce(func_plus, num_list, 1)
# new_result_2: 11
```
