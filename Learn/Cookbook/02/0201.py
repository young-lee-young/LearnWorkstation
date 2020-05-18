# 根据任意多的分割符分割字符串
import re

line = 'asdf fjdk; afed, fjek,asdf, foo'
str_list = re.split(r'[,;\s]\s*', line)
print(str_list)
# str_list = re.split(r'(?:,|;|\s)\s*', line)   # 效果和第一个正则表达式相同


str_list2 = re.split(r'(,|;|\s)\s*', line)  # 包括分割符一起匹配
print(str_list2)


values = str_list2[::2] # 步长是2,如果步长是-1，相当于翻转了列表
print(values)


values = str_list2[1::2] + ['']
print(values)