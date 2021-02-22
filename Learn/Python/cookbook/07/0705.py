# 带有默认参数的函数

def fun_1(a, b=10):
    c = a + b
    return c

print(fun_1(1))
print(fun_1(1, b=20))
print('----------------------')


no_value = object()

def fun_2(a, b=no_value):
    if b is no_value:
        print('b has no value')
    elif b is None:
        print('b is NOne')
    else:
        return a + b

print(fun_2(1))
print(fun_2(1, 20))
print(fun_2(1, None))