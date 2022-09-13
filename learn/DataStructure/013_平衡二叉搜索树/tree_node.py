# -*- coding: utf-8 -*-
# @Time:    2021/1/23 下午11:36
# @Author:  leeyoung
# @File:    tree_node.py
# @Content:


class AVLTreeNode:
    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None
        self.height = 1
