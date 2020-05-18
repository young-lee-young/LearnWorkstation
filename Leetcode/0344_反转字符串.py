# -*- coding: utf-8 -*-
# @Time:    2020/2/27 12:09 PM
# @Author:  leeyoung
# @File:    0344_反转字符串.py
# @Content:


def reverse(str_list):
    begin = 0
    end = len(str_list) - 1

    while begin < end:
        str_list[begin], str_list[end] = str_list[end], str_list[begin]

        begin += 1
        end -= 1
