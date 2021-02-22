# -*- coding: utf-8 -*-
# @Time:    2020/2/26 6:10 PM
# @Author:  leeyoung
# @File:    stack.py
# @Content: 栈

from collections import deque


# 使用双端队列实现栈
class CustomizeStack:
    def __init__(self):
        self.stack = deque()

    def push(self, value):
        return self.stack.append(value)

    def pop(self):
        return self.stack.pop()

    def top(self):
        return self.stack[-1]

    def size(self):
        return len(self.stack)

    def is_empty(self):
        if self.stack:
            return False
        return True
