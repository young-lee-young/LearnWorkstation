# 只接受关键字参数的函数

name = 'liyao'
age = 22
address = 'jilin'
sex = 'boy'


# *后面的都是关键字参数
def fun_1(name, *, age, address, sex):
    print(name, age, address, sex, sep='    ')

fun_1(name, age=age, address=address, sex=sex)


def fun_2(*infos, key=None):
    if key is not None:
        print(infos)
    else:
        print('there is no infos')

fun_2(name, age, address, sex, key=None)
fun_2(name, age, address, sex, key=1)