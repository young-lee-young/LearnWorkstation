# shell通配符做字符串匹配

import os
from fnmatch import fnmatch, fnmatchcase

files = os.listdir('D:/spider')


print(fnmatch(files[1], '*.PNG'))   # windows是不区分大小写的,如果在Linux或者Mac下是False
print(fnmatchcase(files[1], '*.PNG'))   # 精确匹配,windows下也要区分大小写