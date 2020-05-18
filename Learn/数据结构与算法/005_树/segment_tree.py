# -*- coding: utf-8 -*-
# @Time:    2020/3/16 5:49 PM
# @Author:  leeyoung
# @File:    segment_tree.py
# @Content: 线段树

from random import shuffle


class SegmentTree:
    def __init__(self, num_list):
        self.data = num_list
        self.tree = [0] * (4 * len(self.data))
        self._build_segment_tree(0, 0, len(self.data) - 1)

    def _build_segment_tree(self, node_index, begin, end):
        if begin == end:
            self.tree[node_index] = self.data[begin]
            return

        left_index = self._left_child_index(node_index)
        right_index = self._right_child_index(node_index)
        mid = begin + (end - begin) // 2

        self._build_segment_tree(left_index, begin, mid)
        self._build_segment_tree(right_index, mid + 1, end)

        self.tree[node_index] = self.tree[left_index] + self.tree[right_index]

    def _left_child_index(self, index):
        return index * 2 + 1

    def _right_child_index(self, index):
        return index * 2 + 2


num_list = list(range(100))
shuffle(num_list)
segment_tree = SegmentTree(num_list)
print(segment_tree)
