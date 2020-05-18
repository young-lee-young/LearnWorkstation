# -*- coding: utf-8 -*-
# @Time:    2020/2/28 12:38 PM
# @Author:  leeyoung
# @File:    0226.py
# @Content:


def reverse_bin_tree(node):
    if node:
        node.left, node.right = node.right, node.left
        reverse_bin_tree(node.left)
        reverse_bin_tree(node.right)
    return node
