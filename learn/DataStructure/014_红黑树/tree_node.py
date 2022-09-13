# -*- coding: utf-8 -*-
# @Time:    2021/2/3 下午10:39
# @Author:  leeyoung
# @File:    tree_node.py
# @Content:


class RedBlackTreeNode:
    # color默认为True，默认为红色
    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None
        self.color = True
