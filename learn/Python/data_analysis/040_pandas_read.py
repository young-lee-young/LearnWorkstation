

import pandas as pd


# 读取文件,得到DataFrame数据结构,转换为read_csv默认是以逗号为分隔符
pd.read_csv('D:/spider/zhanlang2.csv', encoding='gbk', header=None)    #headers为空,pandas自动为其添加列名


# read_table默认是以制表符为分隔符, sep指定分隔符
pd.read_table('D:/spider/zhanlang2.csv', sep=',', encoding='gbk', names=['name','comment'])


# index_col指定索引列，还可以指定多个索引列: index_col=['年龄','性别']
pd.read_csv('D:/spider/test.csv', encoding='gbk', names=['姓名','年龄','性别'], index_col='年龄')


# 指定读取的行数, nrows=5
pd.read_csv('D:/spider/zhanlang2.csv', encoding='gbk', names=['name', 'comment'], nrows=5)


# 成块读取,chunksize=10,每次读取10条
chunks = pd.read_csv('D:/spider/zhanlang2.csv', encoding='gbk', names=['name', 'comment'], chunksize=10)
print(chunks)
# <pandas.io.parsers.TextFileReader object at 0x000001447752E908>
for chunk in chunks:
        print(chunk)

# 写入csv文件
data_zhanlang2 = pd.read_csv('D:/spider/zhanlang2.csv', encoding='gbk', names=['name', 'comment'])
data_zhanlang2.to_csv('D:/spider/zhanlang2fuben.csv')    #  默认是写入索引和头标题的

data_test = pd.read_csv('D:/spider/test.csv', encoding='gbk', names=['name', 'age', 'sex'])
data_test.to_csv('D:/spider/testfuben.csv', index=False, header=False, columns=['name','age'])    # 可以选择不写入索引和标题, 写入指定的列