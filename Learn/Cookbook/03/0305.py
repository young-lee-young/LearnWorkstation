# 字节串中打包解包

num1 = 58798175117010017516495710119750157


data = num1.to_bytes(16, 'big')
print(data)


num2 = int.from_bytes(data, 'little')
print(num2)