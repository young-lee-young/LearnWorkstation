# -*- coding: utf-8 -*-
# @Time:    2020/3/18 10:31 PM
# @Author:  leeyoung
# @File:    03_全排列.py
# @Content:


def solution(num_list):
    data_len = len(num_list)
    i = 0
    j = 0
    # print(num_list)
    while j < data_len:
        while i < data_len - 1:
            num_list[i], num_list[i + 1] = num_list[i + 1], num_list[i]
            print(num_list)
            i += 1
        j += 1
        i = 0


num_list = [1, 2, 3, 4]
solution(num_list)
