# 在字符串上进行IO操作

import io

text1 = 'hello'
text2 = 'zhangya'

# 操作文本数据
s = io.StringIO()

print(s.write(text1))   # write方法返回的值是写入的长度
print(text2, file=s)
print(s.getvalue()) # 所有写入的数据


s = io.StringIO('liyaozhangya')
print(s.read(4))
print(s.read())
print('---------------------')


# 操作二进制数据
text_b = b'liyaozhangya'
b = io.BytesIO()
print(b.write(text_b))
print(b.getvalue())