# 文件的写入与读取

# 单位换算
# 1G = 1024 M,
# 1M = 2014 KB,
# 1KB = 1024 B(字节,byte),
# 1B = 8 Bit(位)


# 得到一个变量的大小
from sys import getsizeof

name = 'liyao'
print(getsizeof(name))
# 54    得到的单位是字节数


# 文本文件
# 一次性读取,返回字符串,不适合文件过大
with open('D:/spider/maoyan.txt', 'r', encoding='utf-8', errors='ignore') as f:
    text1 = f.read()
    print(text1)

# 读取一行，返回字符串
with open('D:/spider/maoyan.txt', 'r', encoding='utf-8') as f:
    text2 = f.readline()
    print(text2)

# 读取多行,返回列表,每个元素以\n结尾,占用内存比直接读取小
with open('D:/spider/maoyan.txt', 'r', encoding='utf-8') as f:
    text3 = f.readlines()
    print(text3)
print('----------------')


# 文件写入
text_list = ['许巍','汪峰']
# 一次性写入
with open('D:/spider/jingdong.txt','w') as f:
    f.write(text_list[0] + text_list [1])   # 参数需要是一个字符串,如果是列表会报错

# 写入一行
with open('D:/spider/jingdong2.txt','w') as f:
    f.writelines(text_list) # 参数可以是可迭代的对象,循环写入文件,写入后是一行
# 许巍汪峰  # 写入的文件在一行
print('----------------')


# csv文件
import csv

# 读取,可以使用读取文本文件的方式,还可以使用以下方式
with open('D:/spider/zhanlang2.csv') as f:
    csv_file = csv.reader(f) # 返回一个可迭代的对象
    for row in csv_file: # row是一个列表,需要使用下表查找
        print(row[0],row[1])


data_list = [('liyao',22,'男'),('zhangya',21,'女')]
# 写入,使用csv包写入
with open('D:/spider/test.csv','a',newline='') as f:   # 如果不加newline,写出的文件是隔一行一个空格的
    writer = csv.writer(f,dialect='excel')  # dialect定义文件的类型
    for data in data_list:
        writer.writerow(data)

# 写入,使用pandas包写入,正在学习,以后总结
print('---------------')


# json文件
import json

# 读取,可以使用文本文件的读取方法,
with open('D:/spider/json.json','r') as f:
    json_str = json.load(open('D:/spider/json.json'))
    print(json_str)

# 写入文件
json_str = {"users1": [{"name": "liyao", "age": "22"}, {"name": "macheng", "age": "20"}]}
with open('D:/spider/json.json','w',encoding='utf-8') as f:
    json.dump(json_str,f,ensure_ascii=False)    # 如果不加ensure_ascii,写入的是中文的ascii码


import os

path = 'D:/spider'
dirs = os.listdir(path) # 列举出所有文件得到的是一个数组
print(dirs)


isdir = os.path.isdir(path) # 判断是否是路径
isfile = os.path.isfile(path)   # 判断是否是文件
print(isdir, isfile)


import shutil
# 复制文件
shutil.copy('D:/spider/jingdong.txt', 'D:/spider/jingdongcopy.txt')


# 删除文件
os.remove('D:/spider/jingdongcopy.txt')


