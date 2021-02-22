# 迭代出所有的排列和组合

items = ['a', 'b', 'c']


from itertools import permutations  # 排列的模块
for p in permutations(items):   # 三个元素排列
    print(p)
print('-------------------')

for p in permutations(items, 2):    # 两个元素排列
    print(p)
print('-------------------')


from itertools import combinations # 组合的模块
for c in combinations(items, 3):    # 三个元素组合
    print(p)
print('-------------------')

for c in combinations(items, 2):    # 两个元素组合
    print(p)
print('-------------------')

for c in combinations(items, 1):    # 一个元素组合
    print(p)
print('-------------------')


from itertools import combinations_with_replacement
for c in combinations_with_replacement(items, 3):   # 同一个元素可以选择多次
    print(c)