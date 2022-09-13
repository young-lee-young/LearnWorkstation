# 将输出重定向到文件


text1 = 'hello'


with open('D:/spider/test.txt', 'w') as f:
    print(text1, file=f)