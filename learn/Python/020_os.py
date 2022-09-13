import os
import sys


# 获取当前文件的绝对路径
now_path = os.path.abspath(__file__)
print(now_path)


# 获取当前文件的父目录
base_path = os.path.dirname(now_path)
print(base_path)


templates_path = os.path.join(base_path, 'apps', 'templates')
print(templates_path)


print(sys.path.insert(0, base_path))


print(os.environ['TEMP'])