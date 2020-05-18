
# python中的string, tuples, number是不可更改对象, list, dict, set是可更改对象
# python自省, 运行过程中获得对象的类型


# 元组使用
# 列表使用
# 集合使用
# 字典使用


# 列表的使用,列表是按索引查找
# 两个列表和合并为一个列表
num_list1 = [1, 2, 3, 4, 5]
num_list2 = [6, 7, 8, 9, 10]
num_list1.extend(num_list2)
print('和并后的列表:', num_list1)
# 和并后的列表: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]


# 列表切片,切片是重新建立了一个列表,和之前的列表已经没有任何关系
# 切片第一个值是索引是0,第二个值是开区间
num_cut = num_list1[0:3]
print('列表切片:', num_cut)
num_cut[1] = 100
print(num_cut)
print(num_list1)


# 把多个变量分别存到不同的变量
info_list = [('liyao', '10', 'boy'), ('macheng', '20', 'girl'), ('liwei', '30', 'boy')]
for name, age, sex in info_list:
    print('分解为多个变量:', name, age, sex)
print('----------------')


# 集合的使用
sets1 = set()
sets1.add('liyao')  # add是把传入的元素作为一个整体加入到集合中
print('添加元素:', sets1)
# 添加元素: {'liyao'}
sets1.update('zhangya')  # update是把传入的元素分开加到集合中
print('更新元素:', sets1)
# 更新元素: {'h', 'n', 'z', 'y', 'a', 'liyao', 'g'}
sets1.update([1, 2])
print('更新参数是列表的元素:', sets1)
# 更新参数是列表的元素: {'h', 1, 2, 'n', 'z', 'y', 'a', 'liyao', 'g'}
sets1.remove('liyao')  # 删除元素,没有元素会报错
print('删除元素:', sets1)
# 删除元素: {'h', 1, 2, 'n', 'z', 'y', 'a', 'g'}
sets1.pop()  # 随机删除一个元素,没有元素会报错
print('随机删除一个元素:', sets1)
# 随机删除一个元素: {1, 2, 'n', 'h', 'a', 'g', 'z'}
sets1.discard('leeyao')  # 如果有元素就删除,如果没有元素不会报错
print('删除不存在的元素:', sets1)
# 删除不存在的元素: {1, 'g', 2, 'n', 'z', 'y', 'a'}
sets1.clear()
print('清空元素:', sets1)
# 清空元素: set()


# 集合并、交、差操作
sets2 = set('liyao')
sets3 = set('liwei')
sets_bing1 = sets2 | sets3  # 并操作的两种方法
sets_bing2 = sets2.union(sets3)
print('并操作1:', sets_bing1)
print('并操作2:', sets_bing2)
# 并操作1: {'w', 'e', 'l', 'a', 'i', 'y', 'o'}
# 并操作2: {'w', 'e', 'l', 'a', 'i', 'y', 'o'}
sets_jiao1 = sets2 & sets3  # 交操作的两种方法
sets_jiao2 = sets2.intersection(sets3)
print('交操作1:', sets_jiao1)
print('交操作2:', sets_jiao2)
# 交操作1: {'i', 'l'}
# 交操作2: {'i', 'l'}
sets_cha1 = sets2 - sets3  # 差操作的两种方法
sets_cha2 = sets2.difference(sets3)
print('差操作1:', sets_cha1)
print('差操作2:', sets_cha2)
# 差操作1: {'y', 'o', 'a'}
# 差操作2: {'y', 'o', 'a'}
sets_cha3 = sets3 - sets2
sets_cha4 = sets3.difference(sets2)
print('差操作3:', sets_cha3)
print('差操作4:', sets_cha4)
# 差操作3: {'w', 'e'}
# 差操作4: {'w', 'e'}
sets_duichencha1 = sets2 ^ sets3  # 对称差(并集减去交集)操作的两种方法,
sets_duichencha2 = sets2.symmetric_difference(sets3)
print('对称差操作1:', sets_duichencha1)
print('对称差操作2:', sets_duichencha2)
# 对称差操作1: {'w', 'a', 'y', 'e', 'o'}
# 对称差操作2: {'w', 'a', 'y', 'e', 'o'}
print('----------------')


# 字典的使用,字典是按键值对查找,字典是无序的,迭代的是字典的键
info_dict = {
    'name': 'liyao',
    'age': 22,
    'sex': 'boy'
}

keys = info_dict.keys()
print(keys)
values = info_dict.values()
print(values)
items = info_dict.items()
print(items)

if 'name' in info_dict.keys():  # 判断一个键是否在字典里
    name = info_dict['name']
    print(name)

# 获取字典里所有的值,得到的结果可以迭代
values = info_dict.values()
for value in values:
    print(value)

# 删除字典里一个指定元素
info_dict.pop('age')
print(info_dict)

# 随机删除字典里一个元素,在Python3.6中可能是删除最后一个元素
info_dict.popitem()
print(info_dict)

# 清空字典
info_dict.clear()
print(info_dict)

# 字典fromkeys()的用法,用于创建一个新字典
seq = ('name', 'age', 'sex', 'address')
values = ('liyao', 22, 'boy', 'jilin')
new_dict = dict.fromkeys(seq)
print(new_dict)
new_dict2 = dict.fromkeys(seq, values)  # 这种用法不正确,values值不能一一匹配进去
print(new_dict2)
