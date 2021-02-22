# -*- coding: utf-8 -*-
# @Time:    2020/3/15 5:08 PM
# @Author:  leeyoung
# @File:    priority_queue.py
# @Content: 优先队列

from binary_heap import BinaryHeap


class PriorityQueue:
    def __init__(self):
        self.max_heap = BinaryHeap()

    def size(self):
        return self.max_heap.size()

    def is_empty(self):
        return self.max_heap.is_empty()

    def top(self):
        return self.max_heap.top()

    # 入队操作
    def enqueue(self, num):
        self.max_heap.add_node(num)

    # 出队操作
    def dequeue(self):
        return self.max_heap.remove_node()
