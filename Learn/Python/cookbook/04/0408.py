# 跳过可迭代对象的前一部分

from itertools import islice
from itertools import dropwhile

# 如果知道从哪里到哪里不需要,可以用下面的方法
item = ['a', 'b', 'c', 1, 2, 3, 4, 5]
for i in islice(item, 3, None): # 从第三个切到最后
    print(i)
print('-------------------')


# 如果根据特定的特征切片
with open('E:\PycharmWorkstation\Code\Learn\Cookbook\\04\\0401.py', encoding='utf-8') as f:
    for line in f:
        print(line)
print('-------------------')


# dropwhile函数需要两个参数,第一个参数是函数,函数可以返回True和False,第二个参数是一个可以迭代的对象
# dropwhile函数切掉的是开头不符合条件的,在中间的并不能去掉
with open('E:\PycharmWorkstation\Code\Learn\Cookbook\\04\\0401.py', encoding='utf-8') as f:
    for line in dropwhile(lambda line: line.startswith('#'), f):
        print(line)