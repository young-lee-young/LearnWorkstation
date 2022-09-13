# 以键值对的形式迭代

name_list = ['liyao', 'zhangya', 'liuyafei', 'macheng']
for index, name in enumerate(name_list):    # 使用美剧函数
    print(index)
    print(name)
print('---------------------')


# 使用enumerate函数统计文件的行数
with open('E:\PycharmWorkstation\Code\Learn\Cookbook\\04\\0401.py', 'r', encoding='utf-8') as f:
    for lineno, line in enumerate(f, 1):
        print(lineno)
        print(line)
print('---------------------')


# 统计每个此出现的行数
from collections import defaultdict

word_dict = defaultdict(list)

with open('E:\PycharmWorkstation\Code\Learn\Cookbook\\04\\0401.py', 'r', encoding='utf-8') as f:
    lines = f.readlines()

for index, line in enumerate(lines):
    words = [w.strip().lower() for w in line.split()]
    for word in words:
        word_dict[word].append(index)

print(word_dict)