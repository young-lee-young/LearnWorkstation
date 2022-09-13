# filter函数使用

第一个参数方法名，第二个参数是可迭代对象
以可迭代对象每项执行函数，符合条件返回


```python
num_list = range(1, 5)
# num_list: [1, 2, 3, 4]

def func_division(num):
    return num % 2 == 0  # 取余操作

new_num_list = filter(func_division, num_list)
# type(num2): <class 'filter'>

# python2
# new_num_list: [2, 4]
# python3
# <filter object at 0x11042e518>
```
