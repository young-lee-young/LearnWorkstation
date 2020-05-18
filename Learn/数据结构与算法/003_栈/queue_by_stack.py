# -*- coding: utf-8 -*-
# @Time:    2020/2/26 6:11 PM
# @Author:  leeyoung
# @File:    queue.py
# @Content: 使用两个栈实现队列的功能

from customize_stack import CustomizeStack


class CustomizeQueue:
    def __init__(self):
        self.stack_1 = CustomizeStack()
        self.stack_2 = CustomizeStack()

    def append(self, value):
        return self.stack_1.push(value)

    def pop(self):
        if self.stack_2.is_empty():
            if self.stack_1.is_empty():
                return None
            while not self.stack_1.is_empty():
                self.stack_2.push(self.stack_1.pop())
        return self.stack_2.pop()

    def peek(self):
        if not self.stack_2.is_empty():
            return self.stack_2.top()
        while not self.stack_1.is_empty():
            self.stack_1.push(self.stack_1.pop())
        return self.stack_2.top()

    def is_empty(self):
        if self.stack_1.is_empty() and self.stack_2.is_empty():
            return True
        return False
