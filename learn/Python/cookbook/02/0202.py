# 字符串开头和结尾匹配

import os

url = 'http://www.cup.edu.com'

print(url.startswith('https'))
print(url.endswith('.com'))


files = os.listdir('D:/spider')
print(files)
filter_file = [file for file in files if file.endswith(('.png', '.py'))]    # 多个选项,放在一个元组里就好
print(filter_file)

print(any(file.endswith('.ttf') for file in files))
