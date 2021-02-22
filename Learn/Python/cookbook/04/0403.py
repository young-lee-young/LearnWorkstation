
def frange(start, stop, increment): # 生成器函数
    x = start
    while x < stop:
        yield x
        x += increment


for i in frange(0, 5, 1):
    print(i)


print(list(frange(0, 5, 1)))
print(sum(frange(0, 5, 1)))


# 函数中出现了yield就会是函数变成生成器,与普通函数不同,生成器只有在被迭代时才会被执行
def countdown(num):
    print('start {}'.format(num))
    while num > 0:
        yield num
        num -= 1
    print('Done')


c = countdown(3) # 不会执行countdown中的语句
print(c)
print(next(c))
print(next(c))
print(next(c))
print(next(c))  # 这四个next就相当于for循环在迭代,但是for循环已经做了迭代进行完成的处理,next没有错误处理,所以最后会报错
