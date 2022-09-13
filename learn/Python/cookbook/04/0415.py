# 合并多个有序序列,再对整个有序序列进行迭代

import heapq

num_list1 = [1, 3, 5, 6, 7, 9]
num_list2 = [2, 4, 8, 10]


for num in heapq.merge(num_list1, num_list2):
    print(num)


# num_list2不要求函数是有序的,但是如果没有顺序,排列出的也不是有顺序的
num_list2.append(0)
for num in heapq.merge(num_list1, num_list2):
    print(num)