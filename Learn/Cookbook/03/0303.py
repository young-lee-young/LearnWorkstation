# 格式化输出

num = 123.456789


new_num1 = format(num, '0.2f')  # 保留两位小数
print(new_num1)

new_num2 = format(num, '>10.2f')    # 一共10位,两位小数,从前面加空格
print(new_num2)

new_num3 = format(num, '<10.2f')    # 一共10位,两位小数,从后面加空格
print(new_num3)

new_num4 = format(num, '^10.2f')    # 一共10位,两位小数,从两边加空格
print(new_num4)


format_num = 'value is {:0.3f}'.format(num)
print(format_num)
print('------------------')