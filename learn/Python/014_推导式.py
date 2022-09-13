
# 列表推导式
# 分为三部分，导出式，生成式，条件判断式
# 使用[]生成的是列表
num_list = [num for num in range(1, 100) if num % 2 == 0]
print(num_list)


# 使用()生成的是生成器
num_generator = (num for num in range(1, 100))
print(num_generator)
print(type(num_generator))


# 字典推导式
name_list = {'lee', 'ma', 'liu', 'wang', 'zhang','xu'}
new_dict = {key: val for key, val in enumerate(name_list)}
print(new_dict)


# 集合推导式
num_list2 = [1, 2, 3, 4, 5, 1]
new_set = {x for x in num_list2}
print(new_set)
