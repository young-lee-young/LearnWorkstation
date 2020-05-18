# -*- coding: utf-8 -*-
# @Time:    2020/2/26 2:46 PM
# @Author:  leeyoung
# @File:    0007_整数反转.md.py
# @Content:


def reverse_int(num):
    original_num = num
    num = num if num > 0 else num * -1
    return_num = 0

    while num:
        remainder = num % 10
        return_num = return_num * 10 + remainder
        num = num // 10

    return_data = return_num if original_num > 0 else return_num * -1

    if return_data > 2**31 - 1 or return_data < -2**31:
        return 0
    return return_data
