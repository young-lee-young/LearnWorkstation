# -*- coding: utf-8 -*-
# @Time:    2020/3/16 9:03 PM
# @Author:  leeyoung
# @File:    387_第一个唯一字符.py
# @Content: 字符串中第一个唯一字符

from collections import OrderedDict


def first_uniq_char(s):
    letter_dict = OrderedDict()
    for index, item in enumerate(s):
        info = letter_dict.get(item, [0, 0])
        info[0] += 1
        info[1] = index
        letter_dict[item] = info

    for value in letter_dict.values():
        if value[0] == 1:
            return value[1]
    return -1


result = first_uniq_char('leetcode')
print(result)
