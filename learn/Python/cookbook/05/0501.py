# 读写文件

import sys

print(sys.getdefaultencoding()) # 系统默认的编码


text1 = 'liyao'
text2 = 'zhangya'


# Python内部是使用\n来表示换行符的,类Unix平台也是采用\n来表示换行符的,但是windows是采用\r\n两个字符来表示换行符的
# rt模式下,Python读取文件时会把\r\n转换为\n,
# wt模式下,Python写入文件时会用\r\n来表示换行
with open('D:/spider/test.txt', 'wt') as f:
    f.write(text1)
    f.write(text2)  # 两个是都写入,不会产生覆盖

with open('D:/spider/test.txt', 'rt') as f:
    print(f.read())