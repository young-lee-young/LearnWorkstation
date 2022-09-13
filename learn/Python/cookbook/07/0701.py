# 接受任意数量的参数

name = 'liyao'
age = 22
address = 'jilin'
sex = 'boy'

# 接收任意个数位置变量的函数
def fun_1(name, age, *others):
    print(name)
    print(age)
    print(others)   # 其他的变量是以元组的形式存储

fun_1(name, age, address, sex)


# 接收任意数量的关键字参数的函数
def fun_2(name, age, **others):
    print(name)
    print(age)
    print(others)   # 其他的变量是以字典的形式存储

fun_2(name, age, address=address, sex=sex)  # 函数中的关键字参数


# 函数可以接受任意数量的参数和关键字参数
def fun_3(*args, **kwargs):
    print(args)
    print(kwargs)

fun_3(name, age, address=address, sex=sex)