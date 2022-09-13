# 迭代器,生成器,yield,迭代工具的使用


# 迭代器,可以逐一循环出所有的元素,循环的过程叫做迭代
# num_list可以迭代,但不是迭代器
# 所有的元素都存储到内存中
num_list = [1, 2, 3]
for num in num_list:
    print(num)


num_list2 = [num2 * num2 for num2 in range(3)]
for num2 in num_list2:
    print(num2)
print('------------------------')


# 生成器,动态生成值,并没有都加到内存中,第二次执行for循环没有结果
generator = (num3 * num3 for num3 in range(3))
for num3 in generator:
    print(num3)

# 不会执行
for num4 in generator:
    print(num4)
print('-------------------------')


# yield使用
def yield_num():
    num_list3 = [num5 for num5 in range(5)]
    yield from num_list3    # yield后面必须是可迭代的对象

for num6 in yield_num():
    print(num6)


# 迭代工具
import itertools

horses = [1,2,3,4]
races = itertools.permutations(horses)
print(races)
# <itertools.permutations object at 0x000002B3660416D0>
print(list(races))
# [(1, 2, 3, 4),
#  (1, 2, 4, 3),
#  (1, 3, 2, 4),
#  (1, 3, 4, 2),
#  (1, 4, 2, 3),
#  (1, 4, 3, 2),
#  (2, 1, 3, 4),
#  (2, 1, 4, 3),
#  (2, 3, 1, 4),
#  (2, 3, 4, 1),
#  (2, 4, 1, 3),
#  (2, 4, 3, 1),
#  (3, 1, 2, 4),
#  (3, 1, 4, 2),
#  (3, 2, 1, 4),
#  (3, 2, 4, 1),
#  (3, 4, 1, 2),
#  (3, 4, 2, 1),
#  (4, 1, 2, 3),
#  (4, 1, 3, 2),
#  (4, 2, 1, 3),
#  (4, 2, 3, 1),
#  (4, 3, 1, 2),
#  (4, 3, 2, 1)]
