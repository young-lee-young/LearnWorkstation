# 创建不符合系统文件命名规则的文件名

import os
import sys

print(sys.getfilesystemencoding())

with open('D:/spider/test/liyao\xf1o.txt', 'w') as f:
    f.write('something')


print(os.listdir('D:/spider/test/'))
print(os.listdir(b'D:/spider/test'))


'''
这里有问题
'''
print(b'D:/spider/test/')
with open(b'D:/spider/test/liyao\xc3\xb1o.txt', 'r') as f:
    print(f.read())