
# 函数式编程的特点
# 1、把计算视为函数而非指令
# 2、纯函数式编程：不需要变量、没有副作用、测试简单
# 3、支持高阶函数，代码简洁

# Python函数式编程的特点
# Python不是纯函数式编程，允许有变量
# 支持高阶函数，函数也可以作为变量传入
# 支持闭包：有了闭包就能返回函数
# 有限度的支持匿名函数


# 高阶函数：参数可以为函数的函数
def add(x, y, func):
    return func(x) + func(y)


print(add(-5, 9, abs))
print('--------------')


# 变量作用域
# local函数内部作用域，函数内部变量
# enclosing函数内部与内嵌函数之间的作用域，函数和内嵌函数之间的变量
# global全局作用域，整个文件的作用域
# build-in是内置作用域，在Python解释器中，Python内置的变量


# 函数实质与属性
# 1、函数是一个对象
# 2、函数执行完成后内部变量会被回收
# 3、函数属性
# 4、函数返回值


# 闭包：必须有一个内嵌函数，内层函数引用外层函数的变量，返回内层函数
# 作用：封装和代码复用
def add_out(y):
    def add_in(x):
        return x + y
    return add_in


f = add_out(2)
print(type(f))
# <class 'function'>
print(f.__closure__)
# (<cell at 0x000002B7A5850C48: int object at 0x000000006FCCB9F0>,)
print(f(10)) # 实际调用的是add_in，所以不加参数会报错


# 闭包，参数是函数
def func4():
    print('func4被调用')


def func1(func3):
    print('func1被调用')
    def func2():
        print('func2被调用')
        return func3()
    return func2


my_func = func1(func4)
my_func()
print('------------------')



# 装饰器
# 1、装饰器用来装饰函数
# 2、返回一个函数对象
# 3、被装饰函数标识符指向返回的函数对象
# 4、语法糖@demo


# 传进来的是一个函数
def func5(func7):
    print('func5被调用')
    def func6():
        print('func6被调用')
        return func7()
    return func6


# 装饰器实际就是闭包的一个使用，装饰器要在被装饰的函数之前定义，否则会报错
@func5
def func8():
    print('func8被调用')


print(type(func8))
func8()



# 偏函数，把一个多参数的函数转变为少参数的函数
import functools
print(int('100', base=2)) # 二进制转换


def int2(x, base=2):
    return int(x, base)


print(int2('100'))
# 创建一个函数，不需要自己定义
int16 = functools.partial(int, base=16)
print(int16('100'))
