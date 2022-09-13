# 对不存在的文件进行写入操作

text = '对不存在的文件进行写入操作'


# 第一种方法是检查是否存在文件,如果不存在就写入
import os
if not os.path.exists('D:/spider/test.txt'):
    with open('D:/spider/test.txt', 'wt') as f:
        f.write(text)
else:
    print('文件已经存在')


# 第二种方法是使用x模式写入,如果存在文件,会报错
with open('D:/spider/test.txt', 'xt') as f: # 使用x模式写入,如果是二进制使用xb
    f.write(text)