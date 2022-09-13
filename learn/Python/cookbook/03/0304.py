# 二进制、八进制、十六进制

num = 1234


print(bin(num)) # 二进制,前缀是0b
print(oct(-num)) # 八进制,前缀是0o
print(hex(num)) # 十六进制,前缀是0x


print(format(num ,'b')) # 打印出的数据无前缀
print(format(num, 'o'))
print(format(-num, 'x'))


num = int('4d2', 16)    # 把十六进制变成十进制
print(num)