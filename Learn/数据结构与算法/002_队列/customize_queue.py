# -*- coding: utf-8 -*-
# @Time:    2020/2/24 9:45 PM
# @Author:  leeyoung
# @File:    002_队列.py
# @Content: 自定义队列

from collections import deque


class CustomizeQueue:
    def __init__(self):
        # 双端链表
        self.queue = deque()

    def append(self, valkue):
        return self.queue.append(valkue)

    def pop(self):
        return self.queue.popleft()

    def size(self):
        return len(self.queue)

    def is_empty(self):
        if self.queue:
            return False
        return True
