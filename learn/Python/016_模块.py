
# 一个Python文件就是一个模块

# 解决模块同名问题：把相同名称的模块放入到不同的包下面

# 包和普通目录的区别
# 包下面要有__init__.py文件，即使文件是空的，是每一个文件夹下都要有


# 引用模块
try:
    from cStringIO import StringIO
except ImportError:
    from io import StringIo


# Python新旧版本不一样，在旧的版本中可以使用以下导入
from __future__ import unicode_literals # 新的特性名字