# split函数使用
# join函数使用
# replace函数使用

url = 'www.baidu.com'

new_url1 = url.split('.')    # 以.为标志分割，返回的是list
new_url2 = url.split('\n')    #以换行符分割
print(new_url1)
# ['www', 'baidu', 'com']


new_url3 = url.split('.',1)
print(new_url3)
# ['www', 'baidu.com']


new_url4,new_url5 = url.split('.',1)    # 分割一次，保存到两个变量中
print(new_url4,new_url5)
# www baidu.com


# 分割文件路径和文件名
import os

file_path = 'D:/toutiao/wenjian/pic.jpg'
new_url6 = os.path.split(file_path)    #返回的是一个元组，文件路径和文件名被分为两部分
print('分割路径和文件名', new_url6)
# ('D:/toutiao/wenjian', 'pic.jpg')

# 分割路径和后缀
new_url7 = os.path.splitext(file_path)
print('分割路径和后缀', new_url7)
# 分割路径和后缀 ('D:/toutiao/wenjian/pic', '.jpg')


# join函数的使用
str1 = ['liyao','zhang','wang','sun']
str2 = '---'
name = str2.join(str1)
print(name)
# liyao---zhang---wang---sun


# replace函数用法
# 字符串类型是不可变对象,使用replace函数时要重新赋值才会生效
name = 'liyao'
name.replace('i','ee')
print(name)
# liyao    输出结果没变

name = name.replace('i','ee')    # 重新赋值
print(name)
# leeyao    输出结果改变
