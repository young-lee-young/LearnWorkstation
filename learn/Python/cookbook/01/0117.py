# 字典中提取子集

dict_order = {
    'a': 1,
    'b': 2,
    'c': 3,
    'd': 4
}

new_dict = {key: values for key, values in dict_order.items() if values > 2}
print(new_dict)
print('------------------------')