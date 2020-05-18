# 读写压缩的数据文件

import gzip
import bz2

text = '读写压缩的数据文件'
text_b = b'duxieyasuodeshujvwenjian'

# gzip写入与读取
with gzip.open('D:/spider/test.gz', 'wt', encoding='utf-8') as f:
    f.write(text)

with gzip.open('D:/spider/test.gz', 'rt', encoding='utf-8') as f:
    print(f.read())


# bz2写入与读取
with bz2.open('D:/spider/test.bz2', 'wb', compresslevel=5) as f:  # compresslevel指定压缩级别,默认的是最高级别9
    f.write(text_b)

with bz2.open('D:/spider/test.bz2', 'rb') as f:
    print(f.read())