# 打印无法解码的文件名

import os

with open('D:/spider/test/b\udce4d.txt', 'w') as f:
    f.write('hello')

files = os.listdir('D:/spider/test/')
print(files)


def bad_filename(filename):
    return repr(filename)[1:-1] # repr函数将对象转化为功供解释器读取的形式,返回一个对象的string格式


for name in files:
    # print(name)   # 如果直接打印就会报错
    try:
        print(name)
    except UnicodeEncodeError:
        print(bad_filename(name))