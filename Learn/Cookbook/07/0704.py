# 函数返回多个值

# 返回一个数组即可
def return_info():
    name = 'liyao'
    age = 22
    address = 'jilin'
    return name, age, address   # 其实返回的是一个数组

infos = return_info()
print(infos)