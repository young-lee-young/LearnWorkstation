# -*- coding: utf-8 -*-
# @Time:    2020/2/24 9:49 PM
# @Author:  leeyoung
# @File:    003_栈.py
# @Content: 最小值栈（包含获取最小值的方法的栈）

from customize_stack import CustomizeStack


# 两个栈实现最值栈
class MinStack:
    def __init__(self):
        self.stack_1 = CustomizeStack()
        self.stack_2 = CustomizeStack()

    def push(self, value):
        self.stack_1.push(value)

        if self.stack_2.is_empty():
            self.stack_2.push(value)
        else:
            if value <= self.stack_2.top():
                self.stack_2.push(value)

    def pop(self):
        if self.stack_1.is_empty():
            return None

        value = self.stack_1.pop()
        if value == self.stack_2.top():
            self.stack_2.pop()
        return value

    def top(self):
        if self.stack_1.is_empty():
            return None
        return self.stack_1.top()

    def min_value(self):
        if self.stack_1.is_empty():
            return None
        return self.stack_2.top()


def test_func():
    min_stack = MinStack()

    min_stack.push(5)
    min_stack.push(2)

    assert min_stack.top() == 2
    assert min_stack.min_value() == 2

    min_stack.pop()
    assert min_stack.top() == 5
    assert min_stack.min_value() == 5

    min_stack.push(0)
    assert min_stack.top() == 0
    assert min_stack.min_value() == 0

    min_stack.push(1)
    assert min_stack.top() == 1
    assert min_stack.min_value() == 0

    min_stack.pop()
    assert min_stack.top() == 0
    assert min_stack.min_value() == 0
