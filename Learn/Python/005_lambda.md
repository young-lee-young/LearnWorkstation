# lambda表达式

匿名函数

```python
num_list = range(1, 5)

lambda_fun1 = lambda x: x * x
# lambda_fun1: <function <lambda> at 0x107004730>
# type(lambda_fun1): <class 'function'>
new_num_list = map(lambda_fun1, num_list)

lambda_fun2 = lambda x, y: x + y
new_num_list = map(lambda_fun2, num_list, num_list)
```
