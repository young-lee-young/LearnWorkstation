# 字符串变量名插值处理

string = '{name} is {age} years old'
new_string = string.format(name='liyao', age=22)
print(new_string)


name = 'liyao'
age = 23
print(vars())
print(locals())
new_string = string.format_map(vars())  # vars不传值的话,默认在当前环境中搜索
print(new_string)


class Info:
    def __init__(self, name, age):
        self.name = name
        self.age = age

obj = Info('liyao', 24)
new_string = string.format_map(vars(obj))   # vars函数里传入一个类的实例
print(new_string)