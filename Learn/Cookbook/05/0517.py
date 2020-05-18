# 将自己字节数据写入到文本文件

import sys

f = open('D:/spider/test/test.txt', 'w')
print(f.buffer.write(b'hello'))
