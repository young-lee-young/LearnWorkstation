# 字典一键对应多值

from collections import defaultdict

dict_list = defaultdict(list)  # 值是列表,
dict_list['a'].append(1)
print(dict_list)


dict_set = defaultdict(set)  # 值是集合,就是值是不重复的
dict_set['a'].add(1)
print(dict_set)
print('------------------------')