# 检测文件是否存在

import os
import time

print('文件是否存在', os.path.exists('D:/spider/cookies'))

print('是否是文件', os.path.isfile('D:/spider/cookies'))

print('是否是路径', os.path.isdir('D:/spider/cookies'))

print('是否是软连接', os.path.islink('D:/spider/cookieskuaijie'))

print('软连接的真实路径', os.path.realpath('D:/spider/cookieskuaijie'))

print('文件的大小', os.path.getsize('D:/spider/cookies'))

file_time = os.path.getmtime('D:/spider/cookies')
print('文件的修改日期', file_time)
print('格式化时间', time.ctime(file_time))