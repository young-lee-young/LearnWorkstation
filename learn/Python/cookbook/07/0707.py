# 在匿名函数中绑定变量的值

x = 10
a = lambda y: x + y

x = 20
b = lambda y: x + y

print(a(10))
print(b(10))    # 两个结果都是30,lambda函数中的x值是在运行的时候绑定的,不是定义的时候绑定的
print('-----------------------')

# 上面的问题可以这样解决
x = 10
a = lambda y, x=x: x + y

x = 20
b = lambda y, x=x: x + y

print(a(10))    # 结果是20
print(b(10))    # 结果是30
print('-----------------------')


funcs = [lambda x: x+n for n in range(5)]
for f in funcs:
    print(f(0))
print('-----------------------')

funcs = [lambda x, n=n: x+n for n in range(5)]
for f in funcs:
    print(f(0))