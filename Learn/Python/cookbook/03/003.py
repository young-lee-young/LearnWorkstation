# 数值取值

num1 = 123.456
num2 = -123.456


new_num = round(number=num1, ndigits=1)  # 取一位小数,四舍五入
print(new_num)

new_num = round(number=num2, ndigits=2) # 取两位小数
print(new_num)

new_num = round(number=num1, ndigits=-1) # 0是取到个位,-1是取到十位,-2是取到百位
print(new_num)


# 取整操作
num1 = 1.5
num2 = 2.5
new_num1 = round(num1, 0)
new_num2 = round(num2, 0)
print(new_num1, new_num2)   # 都是2.0,是因为取整操作取到该值近的偶数上
print('------------------')