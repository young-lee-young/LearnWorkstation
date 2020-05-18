# -*- coding: utf-8 -*-
# @Time:    2020/2/14 6:22 PM
# @Author:  leeyoung
# @File:    001_两数之和.py
# @Content:


def num_sum(num_list, target):
    num_dict = dict()

    for i in range(len(num_list)):
        temp = target - num_list[i]
        if temp in num_dict:
            return [num_dict[temp], i]
        num_dict[num_list[i]] = i


index_list = num_sum([2, 7, 11, 15], 9)
print(index_list)
