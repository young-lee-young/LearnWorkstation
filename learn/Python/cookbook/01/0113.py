# 通过公共键对字典列表进行排序

from operator import itemgetter

info_dict = [
    {'name': 'liyao', 'age': 25, 'sex': 'boy'},
    {'name': 'zhangya', 'age': 23, 'sex': 'girl'},
    {'name': 'liuyafei', 'age': 22, 'sex': 'boy'},
    {'name': 'macheng', 'age': 21, 'sex': 'boy'}
]


sort_by1 = sorted(info_dict, key=itemgetter('name'))    # 一个键进行排序
print(sort_by1)
# 可以用lambda表达式代替
sort_by1_lambda = sorted(info_dict, key=lambda i: i['name'])
print(sort_by1_lambda)


sort_by2 = sorted(info_dict, key=itemgetter('sex', 'age'))  # 两个键进行排序
print(sort_by2)
sort_by2_lambda = sorted(info_dict, key=lambda i: (i['sex'], i['age']))
print(sort_by2_lambda)
print('------------------------')