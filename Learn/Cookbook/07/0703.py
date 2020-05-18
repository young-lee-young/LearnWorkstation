# 给函数加上参数和返回值的意义

def add(x: int, y: int) -> int: # 只是给参数和返回的值加上了注释,没有实际意义
    return x + y

num_sum = add(1, 3)
print(num_sum)


print(add.__annotations__)