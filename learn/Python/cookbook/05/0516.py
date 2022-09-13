# 为已经打开的文件添加或修改编码方式

import io
from urllib import request

u = request.urlopen('http://www.python.org')    # urlopen是以二进制形式打开的网页
print(u.read())


u = request.urlopen('http://www.python.org')
f = io.TextIOWrapper(u, encoding='utf-8')   # 将以二进制形式打开的文件转换为utf-8编码
print(f.read())


import sys

print(sys.stdout.encoding)
sys.stdout = io.TextIOWrapper(sys.stdout.detach(), encoding='latin-1')
print(sys.stdout.encoding)


f = open('D:/spider/test/test.txt', 'w')
print(f)
# <_io.TextIOWrapper name='D:/spider/test/test.txt' mode='w' encoding='cp936'>
# 文本处理层,负责编码和解码
print(f.buffer)
# <_io.BufferedWriter name='D:/spider/test/test.txt'>
# 缓冲IO层,负责处理二进制数据
print(f.buffer.raw)
# <_io.FileIO name='D:/spider/test/test.txt' mode='wb' closefd=True>
# 原始文件,操作系统底层的文件描述
print('-----------------')


f = io.TextIOWrapper(f.buffer, encoding='latin-1')
print(f)
f.write('hello')    # 会报错,因为上层的文件f已经被销毁





