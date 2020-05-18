# 匿名函数

names = ['David Beazley', 'Brian Jones', 'Raymond Hettinger', 'Ned Batchelder']

new_names = sorted(names, key=lambda names: names.split()[-1].lower())
print(new_names)