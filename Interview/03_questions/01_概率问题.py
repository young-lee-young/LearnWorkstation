# -*- coding: utf-8 -*-
# @Time:    2020/3/18 9:51 AM
# @Author:  leeyoung
# @File:    01.py
# @Content:

# 问题：一个函数，返回 0 和 1，概率为 p 和 1 - p，请你实现一个函数，使得返回 0 和 1 概率一样


def given_func():
    return 1


def customize_func():
    while True:
        first_time = given_func()
        second_time = given_func()

        if first_time == 0 and second_time == 1:
            return 0
        if first_time == 1 and second_time == 0:
            return 1
