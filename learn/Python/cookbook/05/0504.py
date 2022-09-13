# 读写二进制文件

# 常见的二进制文件是视频、音频、图片
with open('D:/spider/background.png', 'rb') as f:
    print(f.read())
print('---------------------')


# 从二进制文件中读取和写入文本,需要进行编码和解码
with open('D:/spider/background.png', 'rb') as f:
    data = f.read()
    print(data.decode('utf-8'))

with open('D:/spider/background.png', 'wb') as f:
    text = 'hello'
    f.write(text.encode('utf-8'))
print('---------------------')


text = 'liyao zhangya'
print(text[0])
for i in text:
    print(i, end=' ')

text_b = b'liyao zhangya'
print(text_b[0])
for i in text_b:
    print(i, end=' ')   # 返回的是代表字节的ASCII码值
print('\n')
print('---------------------')


# 数组和C结构体可以直接写入二进制文件,而不用转换为byte对象
import array

num = array.array('i', [1, 2, 3, 4])
with open('D:/spider/test.bin', 'wb') as f:
    f.write(num)

# 有些对象支持直接将二进制数据读取到他们的底层内存中
a = array.array('i', [0, 0, 0, 0, 0, 0, 0, 0, 0])
with open('D:/spider/test.bin', 'rb') as f:
    f.readinto(a)

print(a)