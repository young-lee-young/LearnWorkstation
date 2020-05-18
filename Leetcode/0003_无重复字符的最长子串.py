# -*- coding: utf-8 -*-
# @Time:    2020/2/14 5:47 PM
# @Author:  leeyoung
# @File:    0003_无重复字符的最长子串.py
# @Content:


def most_length_substr(string):
    new_str = ''
    max_len = 0

    for i in string:
        if i in new_str:
            max_len = len(new_str) if len(new_str) > max_len else max_len
            new_str = new_str.split(i)[1] + i
        else:
            new_str += i
            max_len = len(new_str) if len(new_str) > max_len else max_len

    return max_len


substr_len = most_length_substr("nfpdmpi")
print(substr_len)
