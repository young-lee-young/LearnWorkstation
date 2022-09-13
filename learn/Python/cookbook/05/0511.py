# 处理路径名

import os

path = 'D:/spider/test/test.txt'

print(os.path.basename(path))   # 文件名
print(os.path.dirname(path))    # 路径名

print(os.path.join('D:\\', 'lujing1', 'lujing2', os.path.basename(path)))   # 更改路径名

print(os.path.split(path))  # 分割出文件和路径
print(os.path.splitext(path))   # 分割出文件路径和后缀
