# zip函数的使用

# zip过程
a = [1,2,3]
b = [4,5,6]
c = [7,8,9]
zipped = zip(a,b,c)

print(type(zipped)) # <class 'zip'>

print(zipped)   # <zip object at 0x0000026772F68AC8>

print(dir(zipped))  # dir查看对象的所有属性和方法
#['__class__', '__delattr__', '__dir__', '__doc__', '__eq__', '__format__', '__ge__', '__getattribute__',
# '__gt__', '__hash__', '__init__', '__init_subclass__', '__iter__', '__le__', '__lt__', '__ne__', '__new__',
# '__next__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__']

print(list(zipped))
for i in zipped:    # zip创建了一个迭代器,内容只能被消费一次, 如果上面一句执行, 这句就不会被执行
    print(i)


#unzip过程
zipped2 = [(1, 4, 7), (2, 5, 8), (3, 6, 9)]
unzip= zip(*zipped2)
for j in unzip:
    print(j)

# zip列表转字典
keys = ['name','age','sex']
values = ['lee',22,'boy']
info = dict(zip(keys,values))   # 必须直接转，否则会出错
print(info) # {'name': 'lee', 'age': 22, 'sex': 'boy'}