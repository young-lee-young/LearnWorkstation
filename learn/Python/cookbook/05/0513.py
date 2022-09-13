
import os
import time

path = 'D:/spider/'

dirs = os.listdir(path)
print(dirs)

pathes = [name for name in dirs if os.path.isdir(os.path.join(path, name))] # 全部的路径
print(pathes)

files = [name for name in dirs if os.path.isfile(os.path.join(path, name))] # 全部的文件
print(files)


pyfiles = [name for name in dirs if name.endswith('.py')]
print(pyfiles)


# 文件名的匹配
import glob
from fnmatch import fnmatch

pyfiles = [name for name in dirs if fnmatch(name, '*.py')]
print(pyfiles)
pyfiles = glob.glob(path + '*.py')
print(pyfiles)


# 得到文件名、文件大小、修改时间
name_size_date = [(name, os.path.getsize(name), os.path.getmtime(name)) for name in pyfiles]
for name, size, date in name_size_date:
    print(name, size, time.ctime(date), sep='   ')

files_info = [(name, os.stat(name)) for name in pyfiles]    # 使用stat函数获得文件信息
for name, info in files_info:
    print(name, info.st_size, info.st_mtime)