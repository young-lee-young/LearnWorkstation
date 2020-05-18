# 有序字典
# Python3.6之后字典默认是有序的,就是按照你输入时候的顺序,3.6以前是无序的

from collections import OrderedDict

dict_order = OrderedDict()
dict_order['a'] = 1
dict_order['b'] = 2
dict_order['c'] = 3
dict_order['d'] = 4
print(dict_order)

for i in dict_order.items():
    print(i)
print('------------------------')