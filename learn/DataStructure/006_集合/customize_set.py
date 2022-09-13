# -*- coding: utf-8 -*-
# @Time:    2020/3/15 3:37 PM
# @Author:  leeyoung
# @File:    customize_set.py
# @Content: 自定义集合 - 基于二分搜索树实现

from binary_search_tree import BinarySearchTree


# 性能好于基于链表实现
class CustomizeSet:
    def __init__(self):
        self.binary_search_tree = BinarySearchTree()

    def size(self):
        return self.binary_search_tree.tree_size

    def is_empty(self):
        return self.binary_search_tree.is_empty()

    def contains(self, num):
        # O(h) = O(logn)，h为二分搜索树的高度
        return self.binary_search_tree.contains(num)

    def add(self, num):
        # O(h) = O(logn)
        self.binary_search_tree.add_node(num)

    def remove(self, num):
        # O(h) = O(logn)
        self.binary_search_tree.remove_node(num)


customize_set = CustomizeSet()
customize_set.add(1)
customize_set.add(1)

print(customize_set.size())
print(customize_set.is_empty())
print(customize_set.contains(1))
customize_set.remove(1)
print(customize_set.is_empty())
