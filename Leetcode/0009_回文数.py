# -*- coding: utf-8 -*-
# @Time:    2020/2/27 12:20 PM
# @Author:  leeyoung
# @File:    0009_回文数.py
# @Content:


def is_palindrome(num):
    if num < 0:
        return False

    string = str(num)
    begin, end = 0, len(string) - 1

    while begin < end:
        if string[begin] != string[end]:
            return False
        begin += 1
        end -= 1
    return True
