# -*- coding: utf-8 -*-
# @Time:    2020/2/26 2:35 PM
# @Author:  leeyoung
# @File:    005_堆.py
# @Content: 获取堆topk（需要使用最小堆，每次把最小的元素替换掉）

import heapq
import random


class TopK:
    def __init__(self, num_list, capacity):
        self.min_heap = list()
        self.num_list = num_list
        self.capacity = capacity

    def push(self, num):
        if len(self.min_heap) >= self.capacity:
            heap_top = self.min_heap[0]
            if num > heap_top:
                heapq.heapreplace(self.min_heap, num)
        else:
            heapq.heappush(self.min_heap, num)

    def get_topk(self):
        for num in self.num_list:
            self.push(num)
        return self.min_heap


num_list = list(range(100))
random.shuffle(num_list)
top_k = TopK(num_list, 10)
print(top_k.get_topk())
